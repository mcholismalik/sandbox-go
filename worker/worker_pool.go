package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type JobWrapper struct {
	Func   func(...interface{})
	Params []interface{}
}

type WorkerPool struct {
	MaxGoRoutine int
	MaxEvent     int
	Job          chan JobWrapper
	Wg           *sync.WaitGroup
}

func NewWorkerPool(maxGoRoutine int, maxEvent int) *WorkerPool {
	return &WorkerPool{
		MaxGoRoutine: maxGoRoutine,
		MaxEvent:     maxEvent,
		Job:          make(chan JobWrapper),
		Wg:           &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) AddJob(JobWrapper JobWrapper) {
	wp.Wg.Add(1)
	wp.Job <- JobWrapper
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.MaxGoRoutine; i++ {
		go func(i int) {
			for job := range wp.Job {
				job.Func(job.Params...)
				wp.Wg.Done()
			}
		}(i)
	}
}

func (wp *WorkerPool) Wait() {
	wp.Wg.Wait()
}

func RunWorkerPool() {
	fmt.Println("worker pool - start")

	// benchmark
	defer BenchmarkTime("worker pool", time.Now())
	go NumGoroutine()

	// config
	maxGoRoutine := 3
	maxEvent := 30
	eventResult := make(chan EventResult)
	wp := NewWorkerPool(maxGoRoutine, maxEvent)
	wp.Run()

	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	go func() {
		defer close(wp.Job)
		defer close(eventResult)

		for i := 1; i <= maxEvent; i++ {
			task := JobWrapper{
				Func: func(params ...interface{}) {
					ctxGr := params[0].(context.Context)
					name := params[1].(string)
					i := params[2].(int)

					timeConsume := time.Second * 1
					// if i%3 == 0 {
					// 	timeConsume = time.Second * 10
					// }
					time.Sleep(timeConsume)

					maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, i, timeConsume/time.Second)
					eventResult <- EventResult{
						Name: maskName,
						Err:  ctxGr.Err(),
					}
				},
				Params: []interface{}{ctx, "malik", i},
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

		fmt.Println("worker pool - is error :", result.Err)
	}

	MockUpdateDb("worker pool", successValues, failedValues)

	// DEBUG ONLY - sleep 10s to check goroutine
	fmt.Println("worker pool - sleep 10s")
	time.Sleep(time.Second * 10)

	fmt.Println("worker pool - finish")
}
