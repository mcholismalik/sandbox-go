package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	ctx          context.Context
	maxGoRoutine int
	job          chan func()
	wg           *sync.WaitGroup
}

func NewWorkerPool(ctx context.Context, maxGoRoutine int) *WorkerPool {
	return &WorkerPool{
		ctx:          ctx,
		maxGoRoutine: maxGoRoutine,
		job:          make(chan func()),
		wg:           &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) AddJob(JobWrapper func()) {
	wp.wg.Add(1)
	wp.job <- JobWrapper
}

func (wp *WorkerPool) CloseJob() {
	close(wp.job)
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.maxGoRoutine; i++ {
		go func() {
			for job := range wp.job {
				// skip the rest job, if got ctx error
				if wp.ctx.Err() != nil {
					wp.wg.Done()
					continue
				}

				job()
				wp.wg.Done()
			}
		}()
	}
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

// optional, if need to set as batch / queue
func (wp *WorkerPool) WithBatch(i int) {
	if i%wp.maxGoRoutine == 0 {
		wp.Wait()
	}
}

func RunWorkerPool() {
	fmt.Println("worker pool - start")

	// benchmark
	defer BenchmarkTime("worker pool", time.Now())
	go NumGoroutine()

	// config
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()
	maxGoRoutine := 3
	maxEvent := 30
	eventResult := make(chan EventResult)
	wp := NewWorkerPool(ctx, maxGoRoutine)
	wp.Run()

	go func() {
		defer wp.CloseJob()
		defer close(eventResult)

		for i := 1; i <= maxEvent; i++ {
			_i := i
			name := "malik"
			job := func() {
				timeConsume := time.Second * 1
				// if i%3 == 0 {
				// 	timeConsume = time.Second * 10
				// }
				time.Sleep(timeConsume)

				maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, _i, timeConsume/time.Second)
				eventResult <- EventResult{
					Name: maskName,
					Err:  ctx.Err(),
				}
			}

			wp.AddJob(job)
			// if wanna use batch (old way), put in here
			// wp.WithBatch(i)
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
