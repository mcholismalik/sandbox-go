package generic

import (
	"context"
	"fmt"
	"reflect"
)

type Entity struct {
	ID   int
	Name string
}

type Wrapper struct {
	entity Entity
	ctx    context.Context
}

func generic(any interface{}) {
	params := any.([]interface{})
	wrapper := params[0].(Wrapper)
	fmt.Println(wrapper)
}

func genericSpread(any ...interface{}) {
	fmt.Println("check type", reflect.TypeOf(any))
	fmt.Println("check type 1", reflect.TypeOf(any[0]))
	fmt.Println("check type 2", reflect.TypeOf(any[1]))

}

func New() {
	wrapper := Wrapper{
		entity: Entity{
			ID:   1,
			Name: "cholis",
		},
		ctx: context.Background(),
	}

	params := []interface{}{wrapper}
	generic(params)

	bundle := []interface{}{wrapper, 1, "test"}
	genericSpread(bundle...)
}
