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

	wg := sync.WaitGroup{}
	mut := sync.Mutex{}
	chanWrapper := make(chan string, 100)

	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func() {
				product := Product{
					ID:   "id",
					Name: "name",
				}

				mut.Lock()
				productEvent := product.ToProductEvent()
				mut.Unlock()

				chanWrapper <- productEvent.EventID
			}()
			wg.Done()
		}
	}()

	for i := 1; i <= 100; i++ {
		fmt.Println("result:", <-chanWrapper)
	}

	wg.Wait()

	fmt.Println("finish")
}
