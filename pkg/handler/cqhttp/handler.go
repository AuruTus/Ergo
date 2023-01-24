package cqhttp

import (
	"encoding/json"
	"fmt"

	"github.com/AuruTus/Ergo/pkg/handler"
	ws "github.com/gorilla/websocket"
)

type WSClientHandler struct {
}

func NewWSClientHandler() *WSClientHandler {
	return &WSClientHandler{}
}

var _ handler.WSClientHandler = (*WSClientHandler)(nil)

func (h *WSClientHandler) HandleRead(c *ws.Conn) ([]byte, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("handle read: %w", err)
	}

	return h.msgSwitch(c, msg)
}

func (h *WSClientHandler) HandleWrite(c *ws.Conn, msg []byte) error {
	return c.WriteMessage(ws.TextMessage, msg)
}

func (h *WSClientHandler) msgSwitch(c *ws.Conn, msg []byte) ([]byte, error) {
	// post event
	post := &CommonPostFields{}
	if err := json.Unmarshal(msg, post); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	switch post.PostType {
	case POST_TYPE_MESSAGE, POST_TYPE_MESSAGE_SENT:
		return h.handlePostMessage(c, msg)
	case POST_TYPE_NOTICE:
		fallthrough
	case POST_TYPE_REQUEST:
		return nil, handler.ErrUnimplemented
	case POST_TYPE_META_EVENT:
		return msg, handler.ErrWSControlMsg
	default:
	}

	// api response
	resp := &CommonResponseFields{}
	if err := json.Unmarshal(msg, resp); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	// just ignore response no matter whether it succeeded or failed
	if resp.Status == "ok" || resp.RetCode != 0 {
		return msg, handler.ErrWSResponseMsg
	}

	return msg, nil
}

func (h *WSClientHandler) handlePostMessage(c *ws.Conn, msg []byte) ([]byte, error) {
	message := &PostMessage{}
	if err := json.Unmarshal(msg, message); err != nil {
		return msg, err
	}
	switch message.MessageType {
	case MESSAGE_TYPE_PRIVATE:
		return h.echoHelloPrivate(c, message.UserID)
	case MESSAGE_TYPE_GROUP:
		return h.echoHelloGroup(c, message.GroupID)
	}
	return nil, handler.ErrUnimplemented
}
