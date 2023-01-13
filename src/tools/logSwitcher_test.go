package tools

import (
	"os"
	"reflect"
	"testing"
	"unsafe"

	"github.com/sirupsen/logrus"
)

func TestStringHeader(t *testing.T) {
	str := "Hello, 世界"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	t.Logf("str pointer:%d %d\n", hdr.Data, hdr.Len)
	strSliced := str[:1]
	hdr = (*reflect.StringHeader)(unsafe.Pointer(&strSliced))
	t.Logf("str pointer:%d %d\n", hdr.Data, hdr.Len)

	t.Fail()
}

func TestReflectCompare(t *testing.T) {
	type TestReflectCompareArg struct {
		config      LogConfigs
		logger      *logrus.Logger
		expected    bool
		description string
	}

	textForamatter := new(logrus.TextFormatter)
	hooks := make(logrus.LevelHooks)

	testArgs := []TestReflectCompareArg{
		{LogConfigs{
			"Out":          os.Stdout,
			"Formatter":    textForamatter,
			"Hooks":        hooks,
			"Level":        logrus.InfoLevel,
			"ExitFunc":     os.Exit,
			"ReportCaller": false,
		}, &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    textForamatter,
			Hooks:        hooks,
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		}, true, "TestReflectCompare_1"},

		{LogConfigs{
			"Out":      os.Stdout,
			"Level":    logrus.InfoLevel,
			"ExitFunc": os.Exit,
		}, &logrus.Logger{
			Out:      os.Stderr,
			Level:    logrus.InfoLevel,
			ExitFunc: os.Exit,
		}, false, "TestReflectCompare_2"},
	}

	for i, arg := range testArgs {
		if ok := reflectCompare(arg.config, arg.logger); ok != arg.expected {
			t.Errorf("test %d failed: %s", i, arg.description)
		}
	}
}

func reflectCompare(config LogConfigs, logger *logrus.Logger) bool {
	loggerRef := reflect.ValueOf(logger).Elem()
	for i, len := 0, loggerRef.NumField(); i < len; i++ {
		if c, ok := config[loggerRef.Type().Field(i).Name]; ok &&
			loggerRef.Field(i).Type().Comparable() &&
			c != loggerRef.Field(i).Interface() {
			return false
		}
	}
	return true
}

func TestNewConfiguredLogger(t *testing.T) {
	/* test table */
	type NewConfiguredLoggerTestArg struct {
		configs    LogConfigs
		assertFunc func(LogConfigs, *logrus.Logger) bool
		expected   bool
	}

	testArgs := []NewConfiguredLoggerTestArg{
		{LogConfigs{
			"Out":          os.Stdout,
			"Formatter":    new(logrus.TextFormatter),
			"Hooks":        make(logrus.LevelHooks),
			"Level":        logrus.InfoLevel,
			"ExitFunc":     os.Exit,
			"ReportCaller": false,
		}, reflectCompare, true},

		{LogConfigs{
			"Out":          os.Stderr,
			"dormatter":    new(logrus.TextFormatter),
			"快ooks":        make(logrus.LevelHooks),
			"Level":        logrus.InfoLevel,
			"ExitFunc":     123,
			"ReportCaller": false,
		}, reflectCompare, true},
	}

	/* assertion code */
	for _, arg := range testArgs {
		if logger := NewConfiguredLogger(arg.configs); arg.assertFunc(arg.configs, logger) != arg.expected {
			t.Errorf("invalid config")
		}

	}
}
