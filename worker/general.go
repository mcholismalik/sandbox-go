package worker

import (
	"fmt"
	"runtime"
	"sandbox-go/util"
	"time"
)

type EventResult struct {
	Name string
	Err  error
}

type Person struct {
	ID   int
	Name string
}

func BenchmarkTime(title string, now time.Time) {
	util.TimeTrack(now, title+" - benchmark time")
}

func NumGoroutine() {
	for {
		fmt.Println("num goroutine:", runtime.NumGoroutine())
		time.Sleep(time.Second * 3)
	}
}

func MockUpdateDb(conType string, successValues []string, failedValues []string) {
	msg := fmt.Sprintf("%s - mock update db, successValues:%d failedValues:%d", conType, len(successValues), len(failedValues))
	fmt.Println(msg)
}
