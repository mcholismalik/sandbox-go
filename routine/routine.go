package routine

import "fmt"

type Product struct {
	ID   string
	Name string
}

func New() {
	Cover()
}

func Cover() {
	Loop()
	fmt.Println("Test cover")
}

func Loop() {
	products := []Product{
		{
			ID:   "mock-id-1",
			Name: "mock-name-1",
		},
		{
			ID:   "mock-id-2",
			Name: "mock-name-2",
		},
		{
			ID:   "mock-id-3",
			Name: "mock-name-3",
		},
	}

	for i := range products {
		go func() {
			if i%2 == 0 {
				return
			}
			fmt.Println("Here")
		}()
	}
}
