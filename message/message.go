package message

const (
	StringMessage = 1
	BinMessage = 2
	OfferMessage = 3
	CloseMessage = 4
	HeartBeatMessage = 5
)

type HeartMessage struct {
	Type int `json:"type"`
}

type MessageOffer struct {
	ID int `json:"id"`
	Subscribe    int `json:"subscribe"`
}

type MessageAnswer struct {
	ID int `json:"id"`
	Status   int `json:"status"`
}

type MessageRecive struct {
	Type int `json:"type"`
	Data string `json:"data"`
}

type MessageDispatch struct {
	Type int `json:"type"`
	MsgData []byte `json:"msgData"`
}

type MessageSend struct {
	MsgType int `json:"msgType"`
	Data []byte `json:"data"`
}
