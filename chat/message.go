package chat

const (
	USER =  1
	GROUP =  2
	StringMessage =  3
	BinMessage =  4
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