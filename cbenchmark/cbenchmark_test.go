package cbenchmark

import (
	"fmt"
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

func BenchmarkChannelBool(b *testing.B) {
	b.StopTimer()
	ch := make(chan chan bool)
	go LoopChannelBool(ch)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chB := make(chan bool)
		ch <- chB
		<-chB
	}
}

func BenchmarkChannelString(b *testing.B) {
	b.StopTimer()
	ch := make(chan string)
	go LoopChannelString(ch)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ch <- fmt.Sprintf("abc-%d", i)
	}
}

func BenchmarkChannelStruct(b *testing.B) {
	b.StopTimer()
	ch := make(chan struct{ ID string })
	go LoopChannelStruct(ch)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v := struct{ ID string }{
			ID: fmt.Sprintf("abc-%d", i),
		}
		ch <- v
	}
}
