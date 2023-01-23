package cqhttp

import (
	"encoding/json"
	"fmt"
	"time"

	ws "github.com/gorilla/websocket"
)

/*
	For test
*/
func (h *WSClientHandler) echoHelloPrivate(c *ws.Conn, targetUserID int64) ([]byte, error) {
	return h.sendPrivateInfo(c, targetUserID, "[CQ:face,id=13]Hello,sir!")
}

/*
	API: send_private_msg
*/
func (h *WSClientHandler) sendPrivateInfo(c *ws.Conn, targetUserID int64, msg string) ([]byte, error) {
	action := &ActionSendPrivate{
		CommonActionFields: CommonActionFields{
			Action: API_SEND_PRIVATE_MSG,
			Echo:   fmt.Sprintf("%s-%d", API_SEND_PRIVATE_MSG, time.Now().UnixMicro()),
		},
		Params: ParamsSendPrivateAction{
			UserID:     targetUserID,
			Message:    msg,
			AutoEscape: false,
		},
	}
	return json.Marshal(action)
}
