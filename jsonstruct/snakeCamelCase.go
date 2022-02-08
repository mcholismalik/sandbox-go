package jsonstruct

import (
	"encoding/json"
	"fmt"
)

type HttpImage struct {
	ID       string  `json:"id"`
	FileName *string `json:"file_name"`
}

type GqlImage struct {
	ID string `json:"id"`
}

func New2() {
	httpImage := HttpImage{
		ID: "test bro",
	}

	httpImageMarshal, _ := json.Marshal(httpImage)

	var gqlImage GqlImage
	_ = json.Unmarshal(httpImageMarshal, &gqlImage)

	fmt.Println(gqlImage)

}
