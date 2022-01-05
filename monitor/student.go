package monitor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
	"os"
	"strconv"
)

type Student struct {
	ID     int
	TeacherID []int
	headerData [][]byte
	Send chan message.MessageSend
	Conn *websocket.Conn
}

var fileData [][]byte

var time = 1

func (student *Student) DataRecive() {
	sendData := message.MessageSend{
		message.BinMessage,
		[]byte{},
	}
	for {
		msgType, msg, err := student.Conn.ReadMessage()
		if err != nil {
			//Manager.Unregister <- c
			break
		}

		switch msgType {
		case websocket.TextMessage:
			//Dispatcher.Chat <- msg
		case websocket.BinaryMessage:
			sendData.Data = msg
			if len(student.headerData) < 2{
				student.setHeaderData(msg)
			}else{
				//for _, client := range student.Clients {
				//	client.Send <- sendData
				//}
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

func closeStudent(student *Student) (err error) {
	close(student.Send)
	delete(MonitorManager.Students, student.ID)

	help.Log.Infof("close student %d error", student.ID)

	return nil
}

func (student *Student) closeHandle(code int, text string) error {
	return closeStudent(student)
}

func (student *Student) setHeaderData(data []byte) {
	student.headerData = append(student.headerData, data)
}

func (student *Student) dataDeal(data message.MessageSend){
	student.Send <- data
}

func (student *Student) DataSend() {
	for {
		select {
		case msg, ok := <-student.Send:
			if !ok {
				student.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			switch msg.MsgType {
			case message.StringMessage:
				student.Conn.WriteMessage(websocket.TextMessage, msg.Data)
			case message.BinMessage:
				student.Conn.WriteMessage(websocket.BinaryMessage, msg.Data)
			}
		}
	}
}

func CreateStudent(ID int, conn *websocket.Conn) *Student{
	return &Student{
		ID: ID,
		Conn: conn,
		Send: make(chan message.MessageSend),
	}
}

func StudentConnect(c *gin.Context) {
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

	studentID, err := strconv.Atoi( string(msg) )
	if err != nil {
		help.Log.Infof("student init Atoi err:", err.Error())
		conn.Close()
		return
	}

	student := CreateStudent(studentID, conn)
	MonitorManager.Students[studentID] = student

	student.Conn.SetCloseHandler(student.closeHandle)

	go student.DataRecive()
	go student.DataSend()
}
