package racing

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var index = 0

type Product struct {
	ID   string
	Name string
}

type ProductEvent struct {
	EventID string
	Name    string
	Index   int
}

func (p *Product) ToProductEvent() ProductEvent {
	index = index + 1
	event := ProductEvent{
		EventID: uuid.New().String(),
		Index:   index,
	}
	return event
}

func NewRacing() {
	fmt.Println("start")

	mapUnique := make(map[string]bool)
	numTask := 999999
	wg := sync.WaitGroup{}
	// mut := sync.Mutex{}
	chanWrapper := make(chan string, numTask)

	go func() {
		for i := 1; i < numTask; i++ {
			wg.Add(1)
			go func() {
				product := Product{
					ID:   "id",
					Name: "name",
				}

				// mut.Lock()
				productEvent := product.ToProductEvent()
				fmt.Println(productEvent.EventID)
				if _, ok := mapUnique[productEvent.EventID]; ok {
					panic("gotcha")
				}
				// mut.Unlock()

				if true {
					chanWrapper <- "test return"
					return
				}

				chanWrapper <- productEvent.EventID

			}()
			wg.Done()
		}
	}()

	total := 1
	for i := 1; i < numTask; i++ {
		total = total + 1
		fmt.Println("result:", <-chanWrapper)
		fmt.Println("total:", total)
	}

	wg.Wait()

	fmt.Println("finish")
}
