package tools

import (
	"reflect"
	"runtime"

	"github.com/sirupsen/logrus"
)

type WrappedFunc interface {
	~func()
}

func WithRecover[F WrappedFunc](f F) {
	defer func() {
		if r := recover(); r != nil {
			Log.WithFields(logrus.Fields{"panicRecover": r}).
				Errorf("panic during %s\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
			if err, ok := r.(error); ok {
				Log.Errorf("err in panic: %v\n", err)
			}
		}
	}()
	f()
}

func Go[F WrappedFunc](f F) {
	go WithRecover(f)
}
