package tools

import (
	"fmt"
	"reflect"
	"runtime"
)

type WrappedFunc interface {
	~func()
}

func WithRecover[F WrappedFunc](f F) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"panic during %s recovered: %+v",
				runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
				r,
			)
		}
	}()
	f()
	return
}
