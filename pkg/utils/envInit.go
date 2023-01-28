package utils

import (
	"github.com/AuruTus/Ergo/pkg/utils/configLoader"
	"github.com/AuruTus/Ergo/pkg/utils/logger"
)

func init() {
	configLoader.InitEnv()

	logger.InitLog()
}
