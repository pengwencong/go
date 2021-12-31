package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go/message"
)

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
	HeartTime int
	DataBuff []byte
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
	}()

	//msgFrom := message.MessageFrom{}
	file := false
	video := false

	for {
		messageType, msg, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			break
		}

		if file && messageType == websocket.TextMessage{
			file = false
		}

		if video && messageType == websocket.TextMessage{
			video = false
		}

		switch messageType {
		case websocket.TextMessage:
			//err := json.Unmarshal(msg, &msgFrom)
			//if err != nil{
			//	c.warnF("msgFrom json.Unmarshal err : ", err)
			//	continue
			//}
			Manager.Chat <- msg
			//switch msgFrom.Type {
			//case message.PongMessage:
			//	c.HeartTime -= 1
			//	continue
			//case message.ClientFile:
			//	file = true
			//case message.ClientVideo:
			//	video = true
			//case message.ClientMessage:
			//	Manager.Chat <- msg
			//case message.GroupsMessage:
			//	Manager.Chat <- msg
			//}
		case websocket.BinaryMessage:
			//if file {
			//	err = help.SaveFileFromBinary(c.ID, msgFrom.Content, msg)
			//	file = false
			//}else{
			//	Manager.Monitor <- msg
			//}
		}
	}
}

func (c *Client) textMessageHandle(msgFrom message.MessageFrom){



}

func (c *Client) messageAck(msg string)  {
	msgTo := message.MessageTo{
		From:"",
		Time:"",
		Gid:"0",
		Content:msg,
	}
	toMsg, _ := json.Marshal(msgTo)
	c.Send <- toMsg
}

func (c *Client) warnF(desc string, err error){
	msgTo := message.MessageTo{
		From:"",
		Time:"",
		Gid:"0",
		Content:desc + err.Error(),
	}
	toMsg, _ := json.Marshal(msgTo)
	c.Send <- toMsg
}

func (c *Client) Write() {
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
			fmt.Println("send msg")
			fmt.Println(c.ID, ":", string(message))
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}