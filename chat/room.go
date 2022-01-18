package chat

import (
	"github.com/gorilla/websocket"
	"go/message"
)

type Room struct {
	ID     int
	conn *websocket.Conn
	Send   chan message.MessageSend
}