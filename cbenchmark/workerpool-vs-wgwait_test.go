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

func BenchmarkWgWaitLean(b *testing.B) {
	wg := &sync.WaitGroup{}
	maxGoroutine := 3

	for i := 0; i < b.N; i++ {
		if i%maxGoroutine == 0 {
			wg.Wait()
		}

		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			fmt.Println(WgWaitLean(idx))
		}(i)
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	maxGoroutine := 3
	wp := WorkerPool{
		MaxGoRoutine: maxGoroutine,
		Job:          make(chan JobWrapper),
		Wg:           &sync.WaitGroup{},
	}
	wp.Run()

	for i := 0; i < b.N; i++ {
		taskWrapper := JobWrapper{
			Func: func(params ...interface{}) {
				str := params[0].(string)

				fmt.Println("result:", str)
			},
			Params: []interface{}{fmt.Sprintf("abc-%d", i)},
		}
		wp.AddJob(taskWrapper)
	}
}

func BenchmarkWorkerPoolLean(b *testing.B) {
	maxGoroutine := 3
	wp := WorkerPoolLean{
		MaxGoRoutine: maxGoroutine,
		Job:          make(chan string),
	}
	wp.Run()

	for i := 0; i < b.N; i++ {
		wp.AddJob(fmt.Sprintf("abc-%d", i))
	}
}
