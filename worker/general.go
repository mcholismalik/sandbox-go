package worker

import (
	"fmt"
	"runtime"
	"sandbox-go/util"
	"time"
)

func BenchmarkMemory(title string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("memory consumption: ", title+" - benchmark memory", ByteToMb(m.Alloc), "MiB")
}

func BenchmarkTime(title string, now time.Time) {
	util.TimeTrack(now, title+" - benchmark time")
}

func ByteToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func HighMemoryTask() {
	var overall [][]int
	for i := 0; i < 5; i++ {
		a := make([]int, 0, 999999)
		overall = append(overall, a)
	}
	overall = nil
}
