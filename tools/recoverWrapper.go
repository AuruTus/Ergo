package tools

import (
	"fmt"
	"runtime/debug"
)

type WrappedFunc interface {
	~func()
}

func WithRecover[F WrappedFunc](f F) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %+v\n%s", r, string(debug.Stack()))
		}
	}()
	f()
	return
}
