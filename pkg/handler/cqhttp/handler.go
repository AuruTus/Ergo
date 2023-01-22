package cqhttp

import (
	"github.com/AuruTus/Ergo/pkg/handler"
	ws "github.com/gorilla/websocket"
)

type WSClientHandler struct{}

var _ handler.WSClientHandler = (*WSClientHandler)(nil)

func (h *WSClientHandler) HandleRead(c *ws.Conn) {

}

func (h *WSClientHandler) HandleWrite(c *ws.Conn) {

}
