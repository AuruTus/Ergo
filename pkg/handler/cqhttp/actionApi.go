package cqhttp

import (
	"encoding/json"
	"time"

	"github.com/AuruTus/Ergo/pkg/utils"
	ws "github.com/gorilla/websocket"
)

/*
	For test
*/
func (h *WSClientHandler) echoHelloPrivate(c *ws.Conn, targetUserID int64) ([]byte, error) {
	return h.sendPrivateInfo(c, targetUserID, "[CQ:face,id=13]Hello,sir!")
}

func (h *WSClientHandler) echoHelloGroup(c *ws.Conn, targetGroupID int64) ([]byte, error) {
	return h.sendGroupMsg(c, targetGroupID, "[CQ:face,id=13]Hello,sir!")
}

/*
	api: send_private_msg
*/
func (h *WSClientHandler) sendPrivateInfo(c *ws.Conn, targetUserID int64, msg string) ([]byte, error) {
	action := &ActionSendPrivateMsg{
		CommonActionFields: CommonActionFields{
			Action: API_SEND_PRIVATE_MSG,
			Echo:   utils.KeyGen(API_SEND_PRIVATE_MSG, time.Now().UnixMicro()),
		},
		Params: ParamsSendPrivateMsg{
			UserID:     targetUserID,
			Message:    msg,
			AutoEscape: false,
		},
	}
	return json.Marshal(action)
}

/*
	api: send_group_msg
*/
func (h *WSClientHandler) sendGroupMsg(c *ws.Conn, targetGroupID int64, msg string) ([]byte, error) {
	action := &ActionSendGroupMsg{
		CommonActionFields: CommonActionFields{
			Action: API_SNED_GROUP_MSG,
			Echo:   utils.KeyGen(API_SNED_GROUP_MSG, time.Now().UnixMicro()),
		},
		Params: ParamsSendGroupMsg{
			GroupID:    targetGroupID,
			Message:    msg,
			AutoEscape: false,
		},
	}
	return json.Marshal(action)
}
