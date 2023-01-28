package websocketService

import (
	"fmt"
	"net"
	"net/http"

	"github.com/AuruTus/Ergo/pkg/utils/logger"
)

type WSClientConfig struct {
	HostAddr      net.Addr
	RequestHeader http.Header

	LogConfigs logger.LogConfigs

	WriterBufferSize int
}

// todo: read addr from config
func (c *WSClientConfig) initHostAddr() (err error) {
	c.HostAddr, err = ResolveWSAddrFromSocket("127.0.0.1:8080", "/")
	return
}

// TODO: read header field from config
func (c *WSClientConfig) initRequestHeader() error {
	header := make(http.Header)
	// TODO set token for `Authorization` field
	header.Add("Authorization", "bf7cbe09d71a1bcc373a")
	header.Add("Content-Type", "application/json; charset=utf-8")
	c.RequestHeader = header
	return nil
}

// todo read log config from config
func (c *WSClientConfig) initLogConfig() error {
	c.LogConfigs = make(map[string]any)
	return nil
}

func NewWSClientConfig() (c *WSClientConfig, err error) {
	c = &WSClientConfig{}

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
