package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkerPoolRace struct {
	MaxGoRoutine int
	MaxEvent     int
	Job          chan func()
	Wg           *sync.WaitGroup
}

func NewWorkerPoolRace(maxGoRoutine int, maxEvent int) *WorkerPoolRace {
	return &WorkerPoolRace{
		MaxGoRoutine: maxGoRoutine,
		MaxEvent:     maxEvent,
		Job:          make(chan func()),
		Wg:           &sync.WaitGroup{},
	}
}

func (wp *WorkerPoolRace) AddJob(job func()) {
	wp.Wg.Add(1)
	wp.Job <- job
}

func (wp *WorkerPoolRace) Run() {
	for i := 0; i < wp.MaxGoRoutine; i++ {
		go func(i int) {
			for job := range wp.Job {
				job()
				wp.Wg.Done()
			}
		}(i)
	}
}

func (wp *WorkerPoolRace) Wait() {
	wp.Wg.Wait()
}

func RunWorkerPoolRace() {
	fmt.Println("worker pool race - start")

	// benchmark
	defer BenchmarkTime("worker pool race", time.Now())
	go NumGoroutine()

	// config
	maxGoRoutine := 3
	maxEvent := 30
	eventResult := make(chan EventResult)
	wp := NewWorkerPoolRace(maxGoRoutine, maxEvent)
	wp.Run()

	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	// datas
	persons := []Person{
		{
			ID:   1,
			Name: "Malik",
		},
		{
			ID:   2,
			Name: "Agus Salim",
		},
	}

	go func() {
		defer close(wp.Job)
		defer close(eventResult)

		for _, person := range persons {
			task := func() {
				ctxGr := ctx
				name := person.Name

				timeConsume := time.Second * 1
				// if i%3 == 0 {
				// 	timeConsume = time.Second * 10
				// }
				time.Sleep(timeConsume)

				maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, person.ID, timeConsume/time.Second)
				eventResult <- EventResult{
					Name: maskName,
					Err:  ctxGr.Err(),
				}
			}

			wp.AddJob(task)
		}

		wp.Wait()
	}()

	// update
	failedValues := []string{}
	successValues := []string{}
	for result := range eventResult {
		if result.Err == nil {
			successValues = append(successValues, result.Name)
		} else {
			failedValues = append(failedValues, result.Name)
		}

		fmt.Println("worker pool race - is error :", result.Err)
	}
	MockUpdateDb("worker pool race", successValues, failedValues)

	// DEBUG ONLY - sleep 10s to check goroutine
	fmt.Println("worker pool race - sleep 10s")
	time.Sleep(time.Second * 10)

	fmt.Println("worker pool race - finish")
}
