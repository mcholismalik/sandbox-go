package worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunWgWait() {
	fmt.Println("wg wait - start")

	// benchmark
	defer BenchmarkTime("wg wait", time.Now())
	go NumGoroutine()

	// config
	maxGoroutine := 3
	eventTotal := 30
	eventResult := make(chan EventResult)

	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	go func() {
		// cleansing
		defer close(eventResult)
		wg := sync.WaitGroup{}

		for i := 1; i <= eventTotal; i++ {
			if i%maxGoroutine == 0 {
				wg.Wait()
			}

			wg.Add(1)
			go func(ctxGr context.Context, name string, i int) {
				defer wg.Done()

				timeConsume := time.Second * 1
				// if i%3 == 0 {
				// 	timeConsume = time.Second * 10
				// }
				time.Sleep(timeConsume)

				maskName := fmt.Sprintf(`Mr %s %d, consume %d`, name, i, timeConsume/time.Second)
				eventResult <- EventResult{
					Name: maskName,
					Err:  ctx.Err(),
				}
			}(ctx, "malik", i)
		}

		wg.Wait()
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

		fmt.Println("wg wait - is error :", result.Err)
	}
	MockUpdateDb("wg wait", successValues, failedValues)

	// sleep 10s to check goroutine
	fmt.Println("worker pool - sleep 10s")
	time.Sleep(time.Second * 10)

	fmt.Println("wg wait - finish")
}
