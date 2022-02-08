package s_error

import (
	"fmt"

	"github.com/pkg/errors"
)

func New() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err, ok := errors.Cause(try()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}
	st := err.StackTrace()
	fmt.Printf("%+v", st[0:2]) // top two frames
}

func try() (err error) {
	// defer func() {
	// 	if err != nil {
	// 		err = errors.New("got defer")
	// 		err = errors.Wrap(err, "test")
	// 	}
	// }()

	err = errors.New("not defer")

	return err
}
