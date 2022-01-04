package live

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"strconv"
)

type Room struct {
	ID     int
	Send chan []byte
	Conn *websocket.Conn
	Clients   map[int]*Client
}


func (room *Room) DataRecive() {
	defer func() {
		room.Conn.Close()
	}()

	for {
		msgType, msg, err := room.Conn.ReadMessage()
		if err != nil {
			//Manager.Unregister <- c
			break
		}

		switch msgType {
		case websocket.TextMessage:
			Dispatcher.Chat <- msg
		case websocket.BinaryMessage:
			for _, client := range room.Clients {
				client.Send <- msg
			}
		}
	}
}

func (room *Room) DataSend() {
	defer func() {
		room.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-room.Send:
			if !ok {
				room.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			room.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func CreateRoom(ID int, conn *websocket.Conn) *Room{
	return &Room{
		ID: ID,
		Conn: conn,
		Send: make(chan []byte),
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
