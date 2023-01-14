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

// TODO: read config from file
func (c *WsClientConfig) initRequestHeader() error {
	header := make(http.Header)
	// TODO set token for `Authorization` field
	header.Add("Authorization", "1234")
	header.Add("Content-Type", "application/json; charset=utf-8")
	return nil
}
