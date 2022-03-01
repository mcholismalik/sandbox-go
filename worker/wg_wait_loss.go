package worker

import (
	"fmt"
	"sync"
	"time"
)

func NewWgWaitLoss() {
	fmt.Println("wg wait loss - start")

	for i := 0; i < 10; i++ {
		err := ProcessWgWaitLoss()
		fmt.Println("wg wait loss - result", i, err)

	}

	fmt.Println("wg wait loss - finish")
}

func ProcessWgWaitLoss() error {
	fmt.Println("process wg wait loss - start")

	wg := sync.WaitGroup{}
	processedResult := []string{}

	taskTotal := 9
	counter := 1
	workerTotal := 3
	taskResult := make(chan string, workerTotal)

	for i := 1; i <= taskTotal; i++ {
		eog := counter == taskTotal
		wg.Add(1)
		go func(_i int) {
			defer wg.Done()

			fmt.Println("task sleep 5s")
			time.Sleep(time.Second * 5)

			taskResult <- fmt.Sprintf("task value: %d", _i)
		}(i)

		if i%workerTotal == 0 || eog {
			wg.Wait()
			close(taskResult)

			for t := range taskResult {
				processedResult = append(processedResult, t)
			}
		}

		if !eog {
			taskResult = make(chan string, workerTotal)
		}

		counter++
	}

	for i := range processedResult {
		fmt.Println("task result", processedResult[i])
	}

	fmt.Println("process wait loss - finish")
	return nil
}
