package chat

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)


type Dispatch struct {
	Chat  	   chan []byte
	Mideastream    chan []byte
}


var Dispatcher = &Dispatch{
	Chat:  		make(chan []byte, 10),
	Mideastream:    make(chan []byte),
}

func (dispatch *Dispatch) Start() {
	//go heart()

	for {
		select {
		case chatData := <-dispatch.Chat:
			messageChat := &MessageChat{}
			err := json.Unmarshal(chatData, messageChat)
			if err != nil {

			}

			if client, ok := ChatManager.Clients[messageChat.To]; ok {
				messageSend := MessageSend{
					StringMessage,
					chatData,
				}
				client.Send <- messageSend
			}

		}
	}
}

func createConnect(c *gin.Context) (*websocket.Conn, error) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
		ReadBufferSize:4000,
	}).Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

type chatManager struct {
	Clients map[int]*Client
	Rooms map[int]*Room
}

// Manager define a ws server manager
var ChatManager = &chatManager {
	Clients: make(map[int]*Client, 50),
	Rooms: make(map[int]*Room, 5),
}