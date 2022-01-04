package chat

import (
	"github.com/gorilla/websocket"
)

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
	HeartTime int
	DataBuff []byte
}

func (c *Client) DataRecive() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		_, _, err := c.Socket.ReadMessage()
		if err != nil {
			break
		}
	}
}


func (c *Client) DataSend() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}