package random

import (
	"fmt"
	"math/rand"
)

func New() {
	min := 10000
	max := 99999
	fmt.Println(rand.Intn(max-min) + min)
	fmt.Println(rand.Intn(max-min) + min)
}
