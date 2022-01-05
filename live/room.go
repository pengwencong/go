package live

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
	"strconv"
)

type Room struct {
	ID     int
	Send chan message.MessageSend
	Conn *websocket.Conn
	Clients   map[int]*Client
	headerData [][]byte
}


func (room *Room) DataRecive() {
	defer func() {
		room.Conn.Close()
	}()
	sendData := message.MessageSend{
		message.BinMessage,
		[]byte{},
	}
	for {
		msgType, msg, err := room.Conn.ReadMessage()
		if err != nil {
			//Manager.Unregister <- c
			break
		}

		switch msgType {
		case websocket.TextMessage:
			//Dispatcher.Chat <- msg
		case websocket.BinaryMessage:
			sendData.Data = msg
			if len(room.headerData) < 2{
				room.setHeaderData(msg)
			}else{
				for _, client := range room.Clients {
					client.Send <- sendData
				}
			}
		}
	}
}

func (room *Room) setHeaderData(data []byte) {
	room.headerData = append(room.headerData, data)
}

func (room *Room) dataDeal(data message.MessageSend){
	room.Send <- data
}

func (room *Room) DataSend() {
	defer func() {
		room.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-room.Send:
			if !ok {
				room.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			switch msg.MsgType {
			case message.StringMessage:
				room.Conn.WriteMessage(websocket.TextMessage, msg.Data)
			case message.BinMessage:
				room.Conn.WriteMessage(websocket.BinaryMessage, msg.Data)
			}
		}
	}
}

func CreateRoom(ID int, conn *websocket.Conn) *Room{
	return &Room{
		ID: ID,
		Conn: conn,
		Send: make(chan message.MessageSend),
		Clients:   make(map[int]*Client),
	}
}

func Init(c *gin.Context) {
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

	roomID, err := strconv.Atoi( string(msg) )
	if err != nil {
		help.Log.Infof("room init Atoi err:", err.Error())
		return
	}
	room, ok := LiveManager.Rooms[roomID]
	if ok {
		room.Conn = conn
	} else {
		help.Log.Infof("room init get room Instance err:", err.Error())
		return
	}

	go room.DataRecive()
	go room.DataSend()
}
