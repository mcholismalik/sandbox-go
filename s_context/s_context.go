package s_context

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
		checkMapValues(ctx)
	case <-ctx.Done():
		fmt.Println("halted operation2")

		ctxValueOnly := valueOnlyContext{ctx}

		fmt.Println("check ctx - Done", ctx.Done())
		fmt.Println("check ctxValueOnly - Done", ctxValueOnly.Done())

		checkMapValues(ctxValueOnly)
	}
}

type valueOnlyContext struct{ context.Context }

func (valueOnlyContext) Deadline() (deadline time.Time, ok bool) { return }
func (valueOnlyContext) Done() <-chan struct{}                   { return nil }
func (valueOnlyContext) Err() error                              { return nil }

func checkMapValues(ctx context.Context) {
	for k := range mapValues {
		value := ctx.Value(k).(string)
		fmt.Println("key:", k)
		fmt.Println("value:", value)
	}
}

var mapValues map[string]string

func New() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)

	// set value
	mapValues = make(map[string]string)
	mapValues["x-test-id"] = "mock-positive"
	mapValues["x-app-lang"] = "en"

	for k := range mapValues {
		ctx = context.WithValue(ctx, k, mapValues[k])
	}

	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
