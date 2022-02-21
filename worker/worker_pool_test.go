package worker

import (
	"testing"
)

func BenchmarkNewWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewWorkerPool()
	}
}
