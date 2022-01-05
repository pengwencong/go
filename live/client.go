package live

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
)

// Client is a websocket client
type Client struct {
	ID     int
	Conn *websocket.Conn
	Send   chan message.MessageSend
}

func CreateClient(ID int, conn *websocket.Conn) *Client {
	return &Client{
		ID: ID,
		Conn: conn,
		Send: make(chan message.MessageSend),
	}
}

func ConnectToRoom(c *gin.Context){
	conn, err := createConnect(c)
	if err != nil {
		help.Log.Infof("client connect to room createconnect err:", err.Error())
		return
	}

	conn.SetCloseHandler(closeHandle)

	_, msg, err := conn.ReadMessage()
	if err != nil {
		help.Log.Infof("client connect to room readmsg err:", err.Error())
		conn.Close()
		return
	}

	offer := message.MessageOffer{}
	err = json.Unmarshal(msg, &offer)
	if err != nil {
		help.Log.Infof("client connect to room unmarshal err:", err.Error())
		conn.Close()
		return
	}

	if room, ok := LiveManager.Rooms[offer.Subscribe]; ok {
		client := CreateClient(offer.ID, conn)
		LiveManager.Clients[offer.ID] = client
		room.Clients[offer.ID] = client
		ClientRoomMap.Map[offer.ID] = offer.Subscribe

		go client.DataRecive()
		go client.DataSend()

		msgSend := message.MessageSend{
			message.StringMessage,
			msg,
		}
		msgDispatch := message.MessageDispatch{
			message.OfferMessage,
			msgSend,
		}
		Dispatcher.Chat <- msgDispatch

	} else {
		help.Log.Infof("client connect to room get room err:", err.Error())
		conn.Close()
		return
	}


}

func unregisterClient(client *Client){
	client.Conn.Close()
	close(client.Send)
	delete(LiveManager.Clients, client.ID)
	if roomID, ok := ClientRoomMap.Map[client.ID]; ok {
		delete(LiveManager.Rooms[roomID].Clients, client.ID)
	}
}

func (c *Client) sendHeaderData(headerdata [][]byte){
	for _, val := range headerdata {
		headerData := message.MessageSend{
			message.BinMessage,
			val,
		}
		c.Send <- headerData
	}
}

func closeHandle(code int, text string) error {
	fmt.Println("close")
	fmt.Println(code)
	fmt.Println(text)

	return nil
}

func (c *Client) DataRecive() {
	defer func() {
		//Manager.Unregister <- c
	}()

	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			//Manager.Unregister <- c
			break
		}
		fmt.Println(msg)
		fmt.Println(msgType)
		switch msgType {

		}
	}
}

func (c *Client) DataSend() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			switch msg.MsgType {
			case message.StringMessage:
				c.Conn.WriteMessage(websocket.TextMessage, msg.Data)
			case message.BinMessage:
				c.Conn.WriteMessage(websocket.BinaryMessage, msg.Data)
			}
		}
	}
}