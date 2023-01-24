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
type CommonPostFields struct {
	// event
	Time int64 `json:"time"`
	// qq number of the robot
	SelfID int64 `json:"self_id"`
	// posted event type
	PostType PostType `json:"post_type"`
}

type MessageType string

const (
	MESSAGE_TYPE_PRIVATE MessageType = "private"
	MESSAGE_TYPE_GROUP   MessageType = "group"
)

type Sender struct {
}

/*
	Message events are for things like group and private chat message
*/
type PostMessage struct {
	CommonPostFields
	MessageType MessageType `json:"message_type"`
	SubType     string      `json:"sub_type"`
	MessageID   int32       `json:"message_id"`
	// the sender's qq number
	UserID int64 `json:"user_id"`
	// the group's qq number
	GroupID int64  `json:"group_id"`
	Message string `json:"message"`
	// the message string in cq code format
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
}

/*
	The same as Message, but it's sent from bot.
*/
type PostMessageSent PostMessage

/*
	Notice events are for things like group member notification
*/
type PostNotice struct {
	CommonPostFields
	NoticeType string `json:"notice_type"`
	// TODO complete
}

/*
	Request events are for things like making-friend request
*/
type PostRequest struct {
	CommonPostFields
	RequestType string `json:"request_type"`
	// TODO complete
}

/*
	Meta Event events are for things like cqhttp heartbeat
*/
type PostMetaEvent struct {
	CommonPostFields
	MetaEventType string `json:"meta_event_type"`
	// TODO complete
}
