package tools

import "flag"

/*
	ServiceArgs maintains the enviroment varaibles needed and command
	line aruments passed in.
*/
type ServiceArgs struct {
	ServiceLevel ServiceLevel
}

var EnviromentSettings ServiceArgs

/*
	init function for enviroment args and all global tools
*/
func init() {
	// init enviroment settings
	initServiceLevel()

	// init global tools
	initLog()
}

type ServiceLevel byte

const (
	SERVICE_LEVEL_DEBUG ServiceLevel = iota
	SERVICE_LEVEL_BACKGROUND
)

var serviceLevelMapper = map[string]ServiceLevel{
	"debug":      SERVICE_LEVEL_DEBUG,
	"background": SERVICE_LEVEL_BACKGROUND,
}

/* initServiceLevel get value from "-service-level=" options */
func initServiceLevel() {
	serviceLevel := flag.String("service-level", "debug", "the ServiceLevel enum description arg")
	EnviromentSettings.ServiceLevel = serviceLevelMapper[*serviceLevel]
}
