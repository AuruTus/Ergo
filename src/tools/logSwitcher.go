package tools

import (
	"os"
	"reflect"
	"unicode"

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

func isLogConfigKeyValid(key string) bool {
	isUpperASCIILetter := func(r rune) bool {
		return r < unicode.MaxASCII && unicode.IsUpper(r)
	}

	isValidIdentifierName := func(s string) bool {
		for _, r := range s {
			if r > unicode.MaxASCII || !unicode.IsLetter(r) && r != '_' {
				return false
			}
		}
		return true
	}

	isExportedField := func(key string) bool {
		return isUpperASCIILetter([]rune(key[:1])[0]) &&
			isValidIdentifierName(key)
	}

	return isExportedField(key)
}

/* NewConfiguredLog works for server points */
func NewConfiguredLogger(configs LogConfigs) (logger *logrus.Logger) {
	logger = logSwitcher()
	loggerRef := reflect.ValueOf(logger).Elem()

	for fieldName, config := range configs {
		/*
			NOTE: unicode package will treat excluding letters as Lower or Upper Letters.
			To avoid these undesired rule check, we need define our own ASCII identifier name
			checker
		*/
		if !isLogConfigKeyValid(fieldName) {
			Log.WithFields(logrus.Fields{"fieldName": fieldName}).
				Warnf("Invalid config fieldName\n")
			continue
		}

		configType := reflect.TypeOf(config)
		configVal := reflect.ValueOf(config)

		field := loggerRef.FieldByName(fieldName)
		// if there's invalid config in arguments, neglect it and continue.
		if !field.IsValid() || !field.CanSet() || (!field.CanConvert(configType) && !configType.ConvertibleTo(field.Type())) {
			// the field.Type() equals to the type in the struct definition. It maybe an interface.
			Log.WithFields(logrus.Fields{"fieldName": fieldName, "configType": configType}).
				Warnf("Invalid log config entry\n")
			continue
		}

		field.Set(configVal)
	}
	return
}
