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
	downRate int
	upRate int
	Send   chan message.MessageSend
	Conn *websocket.Conn
}

func CreateTeacher(ID int, conn *websocket.Conn) *Teacher {
	return &Teacher{
		ID: ID,
		Send: make(chan message.MessageSend, 10),
		Conn: conn,
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

		msgDispatch := message.MessageDispatch{
			message.OfferMessage,
			msg,
		}
		Dispatcher.Chat <- msgDispatch

	} else {
		help.Log.Info("teacher connect to room get room err")
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

func (teacher *Teacher) calculateRate(i int) {
	if i % 2 == 0 {
		RateManager.TeacherDownDataLen[teacher.ID] = 0
		RateManager.TeacherUpDataLen[teacher.ID] = 0
	}else{
		teacher.downRate = RateManager.TeacherDownDataLen[teacher.ID] / ( 3 * 1024 )
		teacher.upRate = RateManager.TeacherUpDataLen[teacher.ID] / ( 3 * 1024 )
	}
}

func (teacher *Teacher) DataRecive() {
	for {
		msgType, msg, err := teacher.Conn.ReadMessage()
		if err != nil {
			break
		}

		if _, ok := RateManager.TeacherDownDataLen[teacher.ID]; ok {
			RateManager.TeacherDownDataLen[teacher.ID] += len(msg)
		}else{
			RateManager.TeacherDownDataLen[teacher.ID] = 0
		}

		switch msgType {
		case websocket.TextMessage:
			reciveData := message.MessageRecive{}
			err := json.Unmarshal(msg, &reciveData)
			if err != nil {

			}
			fmt.Printf("%+v\n", reciveData)
		}
	}
}

func (teacher *Teacher) DataSend() {
	for {
		select {
		case msg := <-teacher.Send:

			if _, ok := RateManager.TeacherUpDataLen[teacher.ID]; ok {
				RateManager.TeacherUpDataLen[teacher.ID] += len(msg.Data)
			}else{
				RateManager.TeacherUpDataLen[teacher.ID] = 0
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
