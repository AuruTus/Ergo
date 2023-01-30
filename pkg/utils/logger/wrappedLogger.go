package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

func init() {
	InitLog()
}

/* InitLog should be called in envInit.go's init() */
func InitLog() {
	if Log != nil {
		return
	}
	initLogOnce.Do(func() { Log = logSwitcher() })
}

/* Log is the global Log instance */
var (
	Log         *logrus.Logger
	initLogOnce sync.Once
)

func WithFields(fields logrus.Fields) *logrus.Entry {
	return Log.WithFields(fields)
}

func Infof(format string, args ...any) {
	Log.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	Log.Warnf(format, args...)
}

func Errorf(format string, args ...any) {
	Log.Errorf(format, args...)
}
