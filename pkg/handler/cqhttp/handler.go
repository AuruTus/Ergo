package cqhttp

import (
	"encoding/json"
	"fmt"
	"time"

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

func (h *WSClientHandler) msgSwitch(c *ws.Conn, msg []byte) ([]byte, error) {
	// post evnet
	post := &CommonPostFields{}
	err := json.Unmarshal(msg, post)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	switch post.PostType {
	case POST_TYPE_MESSAGE, POST_TYPE_MESSAGE_SENT:
		message := &PostMessage{}
		json.Unmarshal(msg, message)
		return h.echoHelloPrivate(c, message.UserID)
	case POST_TYPE_NOTICE:
	case POST_TYPE_REQUEST:
	case POST_TYPE_META_EVENT:
		return msg, handler.ErrWSControlMsg
	default:
	}

	// api response
	resp := &CommonResponseFields{}
	err = json.Unmarshal(msg, resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	// just ignore response no matter whether it succeeded or failed
	if resp.Status == "ok" || resp.RetCode != 0 {
		return msg, handler.ErrWSResponseMsg
	}

	return msg, err
}

func (h *WSClientHandler) echoHelloPrivate(c *ws.Conn, targetUserID int64) ([]byte, error) {
	hello := &ActionSendPrivate{
		CommonActionFields: CommonActionFields{
			Action: API_SEND_PRIVATE_MSG,
			Echo:   fmt.Sprintf("%s-%d", API_SEND_PRIVATE_MSG, time.Now().UnixMicro()),
		},
		Params: ParamsSendPrivateAction{
			UserID:     targetUserID,
			Message:    "[CQ:face,id=13]Hello,sir!",
			AutoEscape: false,
		},
	}
	return json.Marshal(hello)
}

func (h *WSClientHandler) HandleWrite(c *ws.Conn, msg []byte) error {
	return c.WriteMessage(ws.TextMessage, msg)
}
