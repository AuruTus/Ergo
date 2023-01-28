package logger

import (
	"os"
	"reflect"

	"github.com/AuruTus/Ergo/pkg/utils/configLoader"
	"github.com/sirupsen/logrus"
)

func logSwitcher() *logrus.Logger {
	switch configLoader.EnviromentSettings.ServiceLevel {
	case configLoader.SERVICE_LEVEL_BACKGROUND:
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
	case configLoader.SERVICE_LEVEL_DEBUG:
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
	loggerRef := reflect.ValueOf(logger).Elem()

	for fieldName, config := range configs {
		// if there's invalid config in arguments, neglect it and continue.
		field := loggerRef.FieldByName(fieldName)

		// Check if the field gotten by config key is valid
		if !field.IsValid() || !field.CanSet() {
			Log.WithFields(logrus.Fields{"fieldName": fieldName}).
				Warnf("Invalid config entry name\n")
			continue
		}

		configType := reflect.TypeOf(config)
		configVal := reflect.ValueOf(config)

		// Check if the config value can be set to the field
		if !field.CanConvert(configType) && !configType.ConvertibleTo(field.Type()) {
			Log.WithFields(logrus.Fields{"fieldName": fieldName, "configType": configType}).
				Warnf("Invalid log config entry value\n")
			continue
		}

		field.Set(configVal)
	}
	return
}
