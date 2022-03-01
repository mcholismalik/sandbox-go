package cbenchmark

import (
	"testing"
)

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
}

func BenchmarkConcatenateBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateBuffer("first", "second")
	}
}

func BenchmarkConcatenateJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateJoin("first", "second")
	}
}

func BenchmarkAppendBigNum(b *testing.B) {
	var overall [][]int
	for i := 0; i < b.N; i++ {
		a := make([]int, 0, 999999)
		overall = append(overall, a)
	}
	overall = nil
}
