package handler

import (
	"errors"

	ws "github.com/gorilla/websocket"
)

type Handler interface {
	WSClientHandler
}

type WSClientHandler interface {
	HandleRead(*ws.Conn) ([]byte, error)
	HandleWrite(*ws.Conn, []byte) error
}

/*
 ErrWSControlMsg is used to denote that there's no need to transport
 message from reader to writer
*/
var ErrWSControlMsg = errors.New("control msg")

var ErrWSResponseMsg = errors.New("response msg")
