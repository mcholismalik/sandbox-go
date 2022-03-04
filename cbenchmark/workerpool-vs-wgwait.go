package cbenchmark

import (
	"fmt"
	"sync"
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

func (wp *WorkerPool) AddJob(taskWrapper JobWrapper) {
	wp.Job <- taskWrapper
}

func (wp *WorkerPool) Run() {
	go func() {
		for i := 0; i < wp.MaxGoRoutine; i++ {
			go func(i int) {
				for task := range wp.Job {
					task.Func(task.Params...)
				}
			}(i)
		}
	}()
}

type WorkerPoolLean struct {
	MaxGoRoutine int
	MaxEvent     int
	Job          chan string
	Wg           *sync.WaitGroup
}

func (wp *WorkerPoolLean) AddJob(taskWrapper string) {
	wp.Job <- taskWrapper
}

func (wp *WorkerPoolLean) Run() {
	go func() {
		for i := 0; i < wp.MaxGoRoutine; i++ {
			go func(i int) {
				for task := range wp.Job {
					func(str string) {
						fmt.Println("result:", str)
					}(task)
				}
			}(i)
		}
	}()
}

func WgWait(ch chan string, wg *sync.WaitGroup, maxGoroutine int) {
	i := 0
	for {
		if i%maxGoroutine == 0 {
			wg.Wait()
		}

		wg.Add(1)
		go func(str string) {
			defer wg.Done()

			fmt.Println("result:", str)
		}(<-ch)

		i++
	}
}

func WgWaitLean(i int) string {
	return fmt.Sprintf("result: %d", i)
}
