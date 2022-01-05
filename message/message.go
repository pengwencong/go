package message

const (
	PongMessage = 0
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

type Test struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

// Message is an object for websocket message which is mapped to json type
type MessageAnswer struct {
	ID int `json:"id"`
	Status   int `json:"status"`
}
