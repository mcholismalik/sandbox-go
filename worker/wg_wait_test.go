package worker

import "testing"

func BenchmarkNewWgWait(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewWgWait()
	}
}
