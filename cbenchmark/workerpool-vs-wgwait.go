package cbenchmark

import (
	"fmt"
	"sync"
)

type TaskWrapper struct {
	Func   func(...interface{})
	Params []interface{}
}

type WorkerPool struct {
	MaxGoRoutine int
	Task         chan TaskWrapper
}

func (wp *WorkerPool) AddTask(taskWrapper TaskWrapper) {
	wp.Task <- taskWrapper
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.MaxGoRoutine; i++ {
		go func(i int) {
			for task := range wp.Task {
				task.Func(task.Params...)
			}
		}(i)
	}
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

			maskStr := fmt.Sprintf(`Mr %s`, str)
			fmt.Println("result:", maskStr)
		}(<-ch)

		i++
	}
}
