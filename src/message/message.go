package message

const (
	PongMessage = 0
	ClientMessage = 1
	GroupsMessage = 2
	ClientFile = 3
	ClientVideo = 4
	Chant_Data = "chat_data"
)

type RedisSetMessage struct {
	Key string `json:"key"`
	Message interface{} `json:"message"`
}

// Message is an object for websocket message which is mapped to json type
type MessageFrom struct {
	From string `json:"from"`
	Time string `json:"time"`
	To    string `json:"to"`
	Type int `json:"type"`
	Content   string `json:"content"`
}

// Message is an object for websocket message which is mapped to json type
type FileMessage struct {
	From    string `json:"to"`
	Time string `json:"time"`
	Name string `json:"name"`
	Content   string `json:"content"`
}

// Message is an object for websocket message which is mapped to json type
type MessageTo struct {
	From    string `json:"to"`
	Time string `json:"time"`
	Gid string `json:"gid"`
	Content   string `json:"content"`
}
