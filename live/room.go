package live

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
	"os"
	"strconv"
)

type Room struct {
	ID     int
	Send chan message.MessageSend
	Conn *websocket.Conn
	Clients   map[int]*Client
	headerData [][]byte
}

var fileData [][]byte

var time = 1

func (room *Room) DataRecive() {
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
			time++
			if time < 10 {
				fileData = append(fileData, msg)
			}
			if time == 10 {
				ff, err := os.Create("./resource/video/room1media.mp4")
				if err != nil {
					fmt.Println("create file err:", err.Error())
				}

				for _, val := range fileData {
					_, err := ff.Write(val)
					if err != nil {
						fmt.Println("file write err: ", err.Error())
					}
				}
				ff.Close()
			}
		}
	}
}

func closeRoom(room *Room) (err error) {
	for _, client := range room.Clients {
		client.Conn.Close()
	}

	close(room.Send)
	delete(LiveManager.Rooms, room.ID)

	help.Log.Infof("close room %d error", room.ID)

	return nil
}

func (room *Room) closeHandle(code int, text string) error {
	return closeRoom(room)
}

func (room *Room) setHeaderData(data []byte) {
	room.headerData = append(room.headerData, data)
}

func (room *Room) dataDeal(data []byte){
	messageSend := message.MessageSend{
		message.StringMessage,
		data,
	}
	room.Send <- messageSend
}

func (room *Room) DataSend() {
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
		conn.Close()
		return
	}

	room := CreateRoom(roomID, conn)
	LiveManager.Rooms[roomID] = room

	room.Conn.SetCloseHandler(room.closeHandle)

	go room.DataRecive()
	go room.DataSend()
}
