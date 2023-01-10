package wsservice

import (
	"net"
	"net/http"
)

type WsClientConfig struct {
	HostAddr      net.Addr
	RequestHeader http.Header
}
