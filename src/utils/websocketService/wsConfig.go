package websocketService

import (
	"fmt"
	"net"
	"net/http"

	"github.com/AuruTus/Ergo/src/tools"
)

type WsClientConfig struct {
	HostAddr      net.Addr
	RequestHeader http.Header

	LogConfigs tools.LogConfigs
}

// todo: read addr from config
func (c *WsClientConfig) initHostAddr() (err error) {
	c.HostAddr, err = ResolveWSAddrFromSocket("127.0.0.1:8080", "/")
	return
}

// TODO: read header field from config
func (c *WsClientConfig) initRequestHeader() error {
	header := make(http.Header)
	// TODO set token for `Authorization` field
	header.Add("Authorization", "bf7cbe09d71a1bcc373a")
	header.Add("Content-Type", "application/json; charset=utf-8")
	c.RequestHeader = header
	return nil
}

// todo read log config from config
func (c *WsClientConfig) initLogConfig() error {
	c.LogConfigs = make(map[string]any)
	return nil
}

func NewWSClientConfig() (c *WsClientConfig, err error) {
	c = &WsClientConfig{}

	if err = c.initHostAddr(); err != nil {
		return nil, fmt.Errorf("initHostAddr: %w", err)
	}
	if err = c.initRequestHeader(); err != nil {
		return nil, fmt.Errorf("initRequestHeader: %w", err)
	}
	if err = c.initLogConfig(); err != nil {
		return nil, fmt.Errorf("initLogConfig: %w", err)
	}

	return
}
