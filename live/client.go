package live

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
	"strconv"
)

// Client is a websocket client
type Client struct {
	ID     int
	Conn *websocket.Conn
	Send   chan []byte
}

func CreateClient(ID int, conn *websocket.Conn) *Client{
	return &Client{
		ID: ID,
		Conn: conn,
		Send: make(chan []byte),
	}
}

func ConnectToRoom(c *gin.Context){
	conn, err := createConnect(c)
	if err != nil {
		help.Log.Infof("room init createConnect err:", err.Error())
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		help.Log.Infof("room init ReadMessage err:", err.Error())
		conn.Close()
		return
	}

	clientId, err := strconv.Atoi( string(msg) )
	if err != nil {
		help.Log.Infof("room init Atoi err:", err.Error())
		return
	}
	client, ok := LiveManager.Clients[clientId]
	if ok {
		client.Conn = conn
	} else {
		help.Log.Infof("room init get room Instance err:", err.Error())
		return
	}

	offer := message.MessageOffer{
		ID: 1,
		Subscribe: 1,
	}
	offerByte, err := json.Marshal(offer)
	if err != nil {
		help.Log.Infof("room init ReadMessage err:", err.Error())
		conn.Close()
		return
	}

	LiveManager.Rooms[1].Send <- offerByte

	go client.DataRecive()
	go client.DataSend()
}

func (c *Client) DataRecive() {
	defer func() {
		//Manager.Unregister <- c
	}()

	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			//Manager.Unregister <- c
			break
		}
	}
}

func (c *Client) DataSend() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}