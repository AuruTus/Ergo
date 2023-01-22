package cqhttp

/*
	NOTE: The event is the json text posted by server. Ref: https://docs.go-cqhttp.org/event/
*/

/*
	the posted event type
*/
type PostType string

const (
	POST_TYPE_MESSAGE      PostType = "message"
	POST_TYPE_MESSAGE_SENT PostType = "message_sent"
	POST_TYPE_REQUEST      PostType = "request"
	POST_TYPE_NOTICE       PostType = "notice"
	POST_TYPE_META_EVENT   PostType = "meta_event"
)

/*
	common fileds for posted event json
*/
type PostCommonFields struct {
	// event
	Time int64 `json:"time"`
	// qq number of the robot
	SelfID int64 `json:"self_id"`
	// posted event type
	PostType PostType `json:"post_type"`
}

/*
	Message events are for things like group and private chat message
*/
type PostMessage struct {
	PostCommonFields
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageID   int32  `json:"message_id"`
	// the sender's qq number
	UserID int64 `json:"user_id"`
	// TODO confirm message type structure
	Message string `json:"message"`
	// the message in cq code format
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	Object     string `json:"object"`
}

/*
	The same as Message, but it's sent from bot.
*/
type MessageSentPosted PostMessage

/*
	Notice events are for things like group member notification
*/
type PostNotice struct {
	PostCommonFields
	NoticeType string `json:"notice_type"`
	// TODO complete
}

/*
	Request events are for things like group member notification
*/
type PostRequest struct {
	PostCommonFields
	RequestType string `json:"request_type"`
	// TODO complete
}

/*
	Meta Event events are for things like cqhttp heatbeat
*/
type PostMetaEvent struct {
	PostCommonFields
	MetaEventType string `json:"meta_event_type"`
	// TODO complete
}
