package chat

const (
	USER = iota
	GROUP
	StringMessage
	BinMessage
)

type MessageChat struct {
	From int `json:"from"`
	To int `json:"to"`
	Type int `json:"type"`
	DataType int `json:"dataType"`
	Data string `json:"data"`
}

type MessageSend struct {
	MsgType int `json:"msgType"`
	Data []byte `json:"data"`
}