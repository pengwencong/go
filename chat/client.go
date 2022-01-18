package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
	"strconv"
)

// Client is a websocket client
type Client struct {
	ID     int
	conn *websocket.Conn
	Send   chan MessageSend
}

func Connect(c *gin.Context) {
	conn, err := createConnect(c)
	if err != nil {
		help.Log.Infof("student init createConnect err:", err.Error())
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		help.Log.Infof("student init ReadMessage err:", err.Error())
		conn.Close()
		return
	}

	clientID, err := strconv.Atoi( string(msg) )
	if err != nil {
		help.Log.Infof("student init Atoi err:", err.Error())
		conn.Close()
		return
	}

	client := CreateClient(clientID, conn)

	ChatManager.Clients[clientID] = client

	go client.DataRecive()
	go client.DataSend()
}

func CreateClient(ID int, conn *websocket.Conn) *Client{
	return &Client{
		ID: ID,
		conn: conn,
		Send: make(chan MessageSend),
	}
}

func closeStudent(student *Client) (err error) {
	close(student.Send)
	help.Log.Infof("close student %d error", student.ID)

	return nil
}

func (c *Client) closeHandle(code int, text string) error {
	return closeStudent(c)
}

func (c *Client) DataRecive() {
	//sendData := message.MessageSend{
	//	message.BinMessage,
	//	[]byte{},
	//}
	for {
		msgType, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		switch msgType {
		case websocket.TextMessage:
			//Dispatcher.Chat <- msgDispatch
		case websocket.BinaryMessage:

		}
	}
}


func (c *Client) DataSend() {
	for {
		select {
		case msg := <-c.Send:

			switch msg.MsgType {
			case message.StringMessage:
				c.conn.WriteMessage(websocket.TextMessage, msg.Data)
			case message.BinMessage:
				c.conn.WriteMessage(websocket.BinaryMessage, msg.Data)
			}

		}
	}
}