package jsonstruct

import (
	"encoding/json"
	"fmt"
)

type Category struct {
	IsKeyProp string `json:"is_key_prop"`
}

func NewUnmarshal() {
	jsonData := `{"is_key_prop":"0"}`

	var category Category

	err := json.Unmarshal([]byte(jsonData), &category)
	if err != nil {
		fmt.Println("Error unmarshal")
	} else {
		fmt.Println("Success unmarshal", category)
	}

}
