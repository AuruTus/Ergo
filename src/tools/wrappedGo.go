package tools

import (
	"reflect"
	"runtime"

	"github.com/sirupsen/logrus"
)

type NestedFunc interface {
	~func()
}

func Go[F NestedFunc](f F) {
	go func() {
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
	}()
}
