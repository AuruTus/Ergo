package cqhttp

type CommonActionFields struct {
	Action string `json:"action"`
	Echo   string `json:"echo,omitempty"`
}

type CommonResponseFields struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Msg     string `json:"msg"`
	Wording string `json:"wording"`
	Echo    string `json:"echo"`
}

const (
	API_SEND_PRIVATE_MSG = "send_private_msg"
)

/*
	api: send_private_msg
*/
type ParamsSendPrivateAction struct {
	UserID     int64  `json:"user_id"`
	GroupID    int64  `json:"group_id,omitempty"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type ActionSendPrivate struct {
	CommonActionFields
	Params ParamsSendPrivateAction `json:"params"`
}

type DataSendPrivate struct {
	MessageID int32 `json:"message_id"`
}

type ResponseSendPrivate struct {
	CommonResponseFields
	Data DataSendPrivate `json:"data"`
}
