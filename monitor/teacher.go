package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
)

// Client is a websocket client
type Teacher struct {
	ID     int
	Conn *websocket.Conn
	Send   chan message.MessageSend
}

func CreateTeacher(ID int, conn *websocket.Conn) *Teacher {
	return &Teacher{
		ID: ID,
		Conn: conn,
		Send: make(chan message.MessageSend),
	}
}

func MonitorStudent(c *gin.Context){
	conn, err := createConnect(c)
	if err != nil {
		help.Log.Infof("teacher connect to room createconnect err:", err.Error())
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		help.Log.Infof("teacher connect to room readmsg err:", err.Error())
		conn.Close()
		return
	}

	offer := message.MessageOffer{}
	err = json.Unmarshal(msg, &offer)
	if err != nil {
		help.Log.Infof("teacher connect to room unmarshal err:", err.Error())
		conn.Close()
		return
	}

	if _, ok := MonitorManager.Students[offer.Subscribe]; ok {
		teacher := CreateTeacher(offer.ID, conn)
		MonitorManager.Teachers[offer.ID] = teacher

		conn.SetCloseHandler(teacher.closeHandle)

		go teacher.DataRecive()
		go teacher.DataSend()

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
		help.Log.Infof("teacher connect to room get room err:", err.Error())
		conn.Close()
		return
	}
}

func unregisterTeacher(teacher *Teacher) (err error) {
	close(teacher.Send)
	delete(MonitorManager.Teachers, teacher.ID)

	help.Log.Infof("unregister teacher %d error", teacher.ID)

	return nil
}

func (teacher *Teacher) sendHeaderData(headerdata [][]byte){
	fmt.Println(len(headerdata))
	for _, val := range headerdata {
		headerData := message.MessageSend{
			message.BinMessage,
			val,
		}
		teacher.Send <- headerData
	}
}

func (teacher *Teacher) closeHandle(code int, text string) error {
	return unregisterTeacher(teacher)
}

func (teacher *Teacher) DataRecive() {
	for {
		msgType, msg, err := teacher.Conn.ReadMessage()
		if err != nil {
			//unregisterTeacher(c)
			break
		}
		fmt.Println(msg)
		fmt.Println(msgType)
		switch msgType {

		}
	}
}

func (teacher *Teacher) DataSend() {
	for {
		select {
		case msg, ok := <-teacher.Send:
			if !ok {
				//unregisterTeacher(c)
				return
			}
			switch msg.MsgType {
			case message.StringMessage:
				teacher.Conn.WriteMessage(websocket.TextMessage, msg.Data)
			case message.BinMessage:
				teacher.Conn.WriteMessage(websocket.BinaryMessage, msg.Data)
			}
		}
	}
}