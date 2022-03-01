package sort

import "fmt"

func NewSort() {
	mapSeq := make(map[int]int)
	mapSeqStr := make(map[string]string)

	mapSeq[2] = 2
	mapSeq[1] = 1
	mapSeq[3] = 3

	mapSeqStr["b"] = "b"
	mapSeqStr["a"] = "a"
	mapSeqStr["c"] = "c"

	fmt.Println("mapSeq")
	for i := range mapSeq {
		fmt.Println("i", i)
		// fmt.Println("v", v)
	}

	fmt.Println("mapSeqStr")
	for k, v := range mapSeqStr {
		fmt.Println("k", k)
		fmt.Println("v", v)
	}
}
