package tools

import (
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

/* initLog should be called in enviromentSettings.go's init() */
func initLog() {
	if Log != nil {
		return
	}
	Log = logSwitcher()
}

/* Log is the global Log instance */
var Log *logrus.Logger

func logSwitcher() *logrus.Logger {
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
	// the debug level logger works as the default.
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

type LogConfigs map[string]any

/* NewConfiguredLog works for server points */
func NewConfiguredLogger(configs LogConfigs) (logger *logrus.Logger) {
	logger = logSwitcher()
	loggerRef := reflect.ValueOf(logger)

	for fieldName, config := range configs {
		configType := reflect.TypeOf(config)
		configVal := reflect.ValueOf(config)

		field := loggerRef.FieldByName(fieldName)
		// if there's invalid config in arguments, neglect it and continue.
		if !field.IsValid() || field.Type() != configType {
			Log.WithFields(map[string]any{"fieldName": fieldName, "configValue": config}).Errorf("Invalid log config entry\n")
			continue
		}

		field.Set(configVal)
	}
	return
}
