package message

const (
	StringMessage = 1
	BinMessage = 2
	OfferMessage = 3
)

type RedisSetMessage struct {
	Key string `json:"key"`
	Message interface{} `json:"message"`
}

// Message is an object for websocket message which is mapped to json type
type MessageOffer struct {
	ID int `json:"id"`
	Subscribe    int `json:"subscribe"`
}

// Message is an object for websocket message which is mapped to json type
type MessageAnswer struct {
	ID int `json:"id"`
	Status   int `json:"status"`
}

type MessageDispatch struct {
	Type int `json:"type"`
	MsgSend MessageSend `json:"msgsend"`
}

type MessageSend struct {
	MsgType int `json:"msgtype"`
	Data []byte `json:"data"`
}
