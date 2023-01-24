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
	API_SNED_GROUP_MSG   = "send_group_msg"
)

/*
	api: send_private_msg
*/
type ParamsSendPrivateMsg struct {
	UserID     int64  `json:"user_id"`
	GroupID    int64  `json:"group_id,omitempty"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type ActionSendPrivateMsg struct {
	CommonActionFields
	Params ParamsSendPrivateMsg `json:"params"`
}

type DataSendPrivateMsg struct {
	MessageID int32 `json:"message_id"`
}

type ResponseSendPrivateMsg struct {
	CommonResponseFields
	Data DataSendPrivateMsg `json:"data"`
}

/*
	api: send_group_msg
*/
type ParamsSendGroupMsg struct {
	GroupID    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

type ActionSendGroupMsg struct {
	CommonActionFields
	Params ParamsSendGroupMsg `json:"params"`
}

type DataSendGroupMsg struct {
	MessageID int32 `json:"message_id"`
}

type ResponseSendGroupMsg struct {
	CommonResponseFields
	Data DataSendGroupMsg `json:"data"`
}
