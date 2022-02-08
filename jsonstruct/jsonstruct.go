package jsonstruct

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Images []ProductImage `json:"images"`
}

type ProductImage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e *Product) ToMarshal() {
	payload, err := json.Marshal(e)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Println(string(payload))
}

func New() {
	product := Product{
		ID:   "abc123",
		Name: "Botol",
		Images: []ProductImage{
			{
				ID:   "def123",
				Name: "Botol 1",
			},
			{
				ID:   "def456",
				Name: "Botol 2",
			},
		},
	}

	product.ToMarshal()

}
