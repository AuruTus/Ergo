package wsservice

import (
	"net"
	"net/http"

	"github.com/AuruTus/Ergo/src/tools"
)

type WsClientConfig struct {
	HostAddr      net.Addr
	RequestHeader http.Header

	LogConfigs tools.LogConfigs
}
