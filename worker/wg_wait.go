package worker

import (
	"fmt"
	"sandbox-go/util"
	"sync"
	"time"
)

func NewWgWait() {
	fmt.Println("wg wait - start")
	defer util.TimeTrack(time.Now(), "wg wait")

	workerTotal := 3
	taskTotal := 9
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

	for i := 0; i < taskTotal; i++ {
		fmt.Println("result wg wait :", <-taskResult)
	}

	fmt.Println("wg wait - finish")
}
