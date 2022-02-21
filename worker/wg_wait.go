package worker

import (
	"fmt"
	"sync"
	"time"
)

func NewWgWait() {
	fmt.Println("wg wait - start")

	// benchmark
	defer BenchmarkMemory("wg wait")
	defer BenchmarkTime("wg wait", time.Now())

	// goroutine checker
	// go func() {
	// 	for {
	// 		fmt.Println("num goroutine:", runtime.NumGoroutine())
	// 		time.Sleep(time.Second * 3)
	// 	}
	// }()

	workerTotal := 3
	taskTotal := 100
	taskResult := make(chan string)

	go func() {
		defer close(taskResult)
		wg := sync.WaitGroup{}

		for i := 1; i <= taskTotal; i++ {
			if i%workerTotal == 0 {
				wg.Wait()
			}

			wg.Add(1)
			go func(name string, i int) {
				defer wg.Done()

				HighMemoryTask()

				timeConsume := time.Second * 1
				if i%3 == 0 {
					timeConsume = time.Second * 10
				}
				time.Sleep(timeConsume)

				maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, i, timeConsume/time.Second)
				taskResult <- maskName
			}("malik", i)
		}

		wg.Wait()
	}()

	for result := range taskResult {
		fmt.Println("wg wait - result :", result)
	}

	fmt.Println("wg wait - finish")
}
