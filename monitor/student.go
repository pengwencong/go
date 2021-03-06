package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	"go/message"
	"strconv"
)

type Student struct {
	ID     int
	framerType FramerType
	downRate int
	upRate int
	headerData [][]byte
	Send chan message.MessageSend
	Conn *websocket.Conn
}

//var fileData [][]byte

//var time = 1

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

		RateManager.StudentDownDataLen[student.ID] += len(msg)

		switch msgType {
		case websocket.TextMessage:
			//Dispatcher.Chat <- msgDispatch
		case websocket.BinaryMessage:
			sendData.Data = msg
			if len(student.headerData) < 2 {
				student.setHeaderData(msg)
			} else {
				for _, teacher := range MonitorManager.Teachers {
					teacher.Send <- sendData
				}
			}
			//time++
			//if time < 10 {
			//	fileData = append(fileData, msg)
			//}
			//if time == 10 {
			//	ff, err := os.Create("./resource/video/room1media.mp4")
			//	if err != nil {
			//		fmt.Println("create file err:", err.Error())
			//	}
			//
			//	for _, val := range fileData {
			//		_, err := ff.Write(val)
			//		if err != nil {
			//			fmt.Println("file write err: ", err.Error())
			//		}
			//	}
			//	ff.Close()
			//}
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

func (student *Student) dataDeal(data []byte){
	messageSend := message.MessageSend{
		message.StringMessage,
		data,
	}
	student.Send <- messageSend
}

func (student *Student) DataSend() {
	for {
		select {
		case msg := <-student.Send:

			RateManager.StudentUpDataLen[student.ID] += len(msg.Data)

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
		framerType: Framer30,
		Conn: conn,
		Send: make(chan message.MessageSend),
	}
}

func (student *Student) calculateRate(i int)  int {
	if i % 2 == 0 {
		RateManager.StudentDownDataLen[student.ID] = 0
		RateManager.StudentUpDataLen[student.ID] = 0
	} else {
		student.downRate = RateManager.StudentDownDataLen[student.ID] / ( 3 * 1024 )
		student.upRate = RateManager.StudentUpDataLen[student.ID] / ( 3 * 1024 )
		fmt.Println(student.downRate)
		fmt.Println(RateManager.StudentDownDataLen[student.ID])
	}

	return student.downRate
}

func (student *Student) adjustMediaFramer(i int) {
	rate := student.calculateRate(i)
	diff := W110H150f30 - rate
	if diff <= 0 {
		return
	}

	framerMsg := message.MessageAdjust{
		message.AdjustFramer,
		int(Framer30),
	}
	msg := message.MessageSend{
		message.StringMessage,
		[]byte{},
	}

	//var once sync.Once
	//onceBody := func() {
		//time.Sleep(time.Second * 4)
		//
		//framerMsg.FramerType = W110H150f20
		//framerMsgByte, _ := json.Marshal(framerMsg)
		//msg.Data = framerMsgByte
		//
		//student.Send <- msg
	//}
	//once.Do(onceBody)

	switch student.framerType {
	case Framer30:
		if 4 < diff && diff < 10 {
			framerMsg.FramerType = W110H150f20
			framerMsgByte, _ := json.Marshal(framerMsg)
			msg.Data = framerMsgByte

			student.Send <- msg
		}
	case Framer20:
		if 14 < diff && diff < 20 {
			framerMsg.FramerType = W110H150f10
			framerMsgByte, _ := json.Marshal(framerMsg)
			msg.Data = framerMsgByte

			student.Send <- msg
		}
	case Framer10:

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
