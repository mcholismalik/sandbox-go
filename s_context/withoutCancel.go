package s_context

import (
	"context"
	"fmt"
	"time"
)

type ValueOnlyContext struct{ context.Context }

func (ValueOnlyContext) Deadline() (deadline time.Time, ok bool) { return }
func (ValueOnlyContext) Done() <-chan struct{}                   { return nil }
func (ValueOnlyContext) Err() error                              { return nil }

func NewWithoutCancel() {
	fmt.Println("withoutCancel -> start")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	if true {
		cancel()
	}

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
		}
	}(ctx)

	fmt.Println("withoutCancel -> finish")
}
