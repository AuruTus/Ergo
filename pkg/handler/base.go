package handler

import (
	ws "github.com/gorilla/websocket"
)

type Handler interface {
	WSClientHandler
}

type WSClientHandler interface {
	HandleRead(*ws.Conn)
	HandleWrite(*ws.Conn)
}
