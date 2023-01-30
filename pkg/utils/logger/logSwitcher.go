package logger

import (
	"flag"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

type ServiceLevel byte

const (
	SERVICE_LEVEL_DEBUG ServiceLevel = iota
	SERVICE_LEVEL_BACKGROUND
)

var serviceLevelMapper = map[string]ServiceLevel{
	"debug":      SERVICE_LEVEL_DEBUG,
	"background": SERVICE_LEVEL_BACKGROUND,
}

func logSwitcher() *logrus.Logger {

	serviceLevelEnv := flag.String("service-level", "debug", "the ServiceLevel enum description arg")
	serviceLevel := serviceLevelMapper[*serviceLevelEnv]
	switch serviceLevel {
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
