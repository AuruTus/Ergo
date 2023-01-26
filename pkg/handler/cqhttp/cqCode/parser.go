package cqMessage

import "strings"

type CQMessageType string

const (
	CQ_MESSAGE_TYPE_TEXT   CQMessageType = "text"
	CQ_MESSAGE_TYPE_FACE   CQMessageType = "face"
	CQ_MESSAGE_TYPE_RECORD CQMessageType = "record"
	CQ_MESSAGE_TYPE_VIDEO  CQMessageType = "video"
	CQ_MESSAGE_TYPE_AT     CQMessageType = "at"
	CQ_MESSAGE_TYPE_RPS    CQMessageType = "rps"
	CQ_MESSAGE_TYPE_DICE   CQMessageType = "dice"
	CQ_MESSAGE_TYPE_MUSIC  CQMessageType = "music"
	CQ_MESSAGE_TYPE_IMAGE  CQMessageType = "image"
	CQ_MESSAGE_TYPE_REPLY  CQMessageType = "reply"
)

type CQCodeNode struct {
	Type string
	Data map[string]string
}

func ParseCQCode(code string) *CQCodeNode {
	cq := new(CQCodeNode)
	curr, code, _ := strings.Cut(code, ",")
	cq.Type = curr[4:]
	for range code {
		curr, code, ok := strings.Cut(code, ",")
		switch ok {
		case true:
			k, v, _ := strings.Cut(curr, "=")
			cq.Data[k] = v
		default:
			k, v, _ := strings.Cut(code, "=")
			cq.Data[k] = v
		}
	}
	return cq
}
