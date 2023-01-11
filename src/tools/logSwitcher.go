package tools

import (
	"os"

	"github.com/sirupsen/logrus"
)

/* initLog should be called in enviromentSettings.go's init() */
func initLog() {
	if Log != nil {
		return
	}
	Log = LogSwitcher()
}

var Log *logrus.Logger

func LogSwitcher() *logrus.Logger {
	switch EnviromentSettings.ServiceLevel {
	case SERVICE_LEVEL_BACKGROUND:
		return &logrus.Logger{
			// TODO change background level config (the output file!)
			Out:          os.Stderr,
			Formatter:    new(logrus.TextFormatter),
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		}
	/* the debug level logger works as the default.  */
	case SERVICE_LEVEL_DEBUG:
		fallthrough
	default:
		return &logrus.Logger{
			Out:          os.Stderr,
			Formatter:    new(logrus.TextFormatter),
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		}
	}
}
