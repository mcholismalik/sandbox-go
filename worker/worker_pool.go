package worker

import (
	"fmt"
	"sandbox-go/util"
	"sync"
	"time"
)

type TaskWrapper struct {
	Func  func(...interface{})
	Param []interface{}
}

type WorkerPool struct {
	TotalWorker int
	Wg          *sync.WaitGroup
	Task        chan TaskWrapper
}

func (wp *WorkerPool) AddTask(taskWrapper TaskWrapper) {
	wp.Task <- taskWrapper
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.TotalWorker; i++ {
		go func(i int) {
			for task := range wp.Task {
				task.Func(task.Param...)
				wp.Wg.Done()
			}
		}(i)
	}
}

func NewWorkerPool() {
	fmt.Println("worker pool - start")
	defer util.TimeTrack(time.Now(), "worker pool")

	// worker
	wg := sync.WaitGroup{}
	workerTotal := 3
	worker := WorkerPool{
		Wg:          &wg,
		TotalWorker: workerTotal,
		Task:        make(chan TaskWrapper),
	}
	worker.Run()

	// task
	taskTotal := 9
	taskResult := make(chan string, taskTotal)
	wg.Add(taskTotal)

	for i := 1; i <= taskTotal; i++ {
		task := TaskWrapper{
			Func: func(params ...interface{}) {
				name := params[0].(string)
				i := params[1].(int)

				timeConsume := time.Second * 1
				if i%3 == 0 {
					timeConsume = time.Second * 10
				}
				time.Sleep(timeConsume)

				maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, i, timeConsume/time.Second)
				taskResult <- maskName
			},
			Param: []interface{}{"malik", i},
		}
		worker.AddTask(task)
	}

	for i := 0; i < taskTotal; i++ {
		fmt.Println("result worker pool :", <-taskResult)
	}

	close(taskResult)
	close(worker.Task)

	fmt.Println("worker pool - finish")
}
