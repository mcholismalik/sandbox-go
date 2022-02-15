package worker

import (
	"fmt"
	"time"
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

func NewWorkerPool() {
	fmt.Println("worker pool - start")

	// benchmark
	defer BenchmarkMemory("worker pool")
	defer BenchmarkTime("worker pool", time.Now())

	// goroutine checker
	// go func() {
	// 	for {
	// 		fmt.Println("num goroutine:", runtime.NumGoroutine())
	// 		time.Sleep(time.Second * 3)
	// 	}
	// }()

	// worker
	maxGoRoutine := 3
	worker := WorkerPool{
		MaxGoRoutine: maxGoRoutine,
		Task:         make(chan TaskWrapper),
	}
	worker.Run()

	// task
	maxTask := 100
	taskResult := make(chan string, maxTask)

	// defer
	defer func() {
		close(taskResult)
		close(worker.Task)
	}()

	go func() {
		for i := 1; i <= maxTask; i++ {
			task := TaskWrapper{
				Func: func(params ...interface{}) {
					name := params[0].(string)
					i := params[1].(int)

					HighMemoryTask()

					timeConsume := time.Second * 1
					// if i%3 == 0 {
					// 	timeConsume = time.Second * 10
					// }
					time.Sleep(timeConsume)

					maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, i, timeConsume/time.Second)
					taskResult <- maskName
				},
				Params: []interface{}{"malik", i},
			}
			worker.AddTask(task)
		}
	}()

	for i := 0; i < maxTask; i++ {
		fmt.Println("worker pool - result :", <-taskResult)
	}

	fmt.Println("worker pool - finish")
}
