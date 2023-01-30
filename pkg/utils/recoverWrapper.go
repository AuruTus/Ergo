package utils

import (
	"fmt"
	"runtime/debug"
)

type WrappedPanic struct {
	r         any
	stackInfo []byte
}

var _ error = (*WrappedPanic)(nil)

func (e *WrappedPanic) Error() string {
	return fmt.Sprintf("panic recovered: %+v\n%s", e.r, e.stackInfo)
}

type WrappedFunc interface {
	~func()
}

func WithRecover[F WrappedFunc](f F) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &WrappedPanic{r: r, stackInfo: debug.Stack()}
		}
	}()
	f()
	return
}
