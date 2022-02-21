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
