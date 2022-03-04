package cbenchmark

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkWgWait(b *testing.B) {
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	maxGoroutine := 3
	go WgWait(ch, wg, maxGoroutine)

	for i := 0; i < b.N; i++ {
		ch <- fmt.Sprintf("abc-%d", i)
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	maxGoroutine := 3
	wp := WorkerPool{
		MaxGoRoutine: maxGoroutine,
		Task:         make(chan TaskWrapper),
	}
	wp.Run()

	for i := 0; i < b.N; i++ {
		taskWrapper := TaskWrapper{
			Func: func(params ...interface{}) {
				str := params[0].(string)

				fmt.Println("result:", str)
			},
			Params: []interface{}{fmt.Sprintf("abc-%d", i)},
		}
		wp.AddTask(taskWrapper)
	}
}
