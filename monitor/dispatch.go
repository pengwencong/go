package monitor

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/message"
	"net/http"
	"time"
)

// ClientManager is a websocket manager
type Dispatch struct {
	Chat  	   chan message.MessageDispatch
	Mideastream    chan []byte
	NetRateMonitor chan *Student
}


var Dispatcher = &Dispatch{
	Chat:  		make(chan message.MessageDispatch),
	Mideastream:    make(chan []byte),
	NetRateMonitor: make(chan *Student, 50),
}


func (dispatch *Dispatch) Start() {
	go heart()

	for {
		select {
		case msgDispatch := <-dispatch.Chat:
			switch msgDispatch.Type {
			case message.OfferMessage:
				offer := message.MessageOffer{}
				json.Unmarshal(msgDispatch.MsgData, &offer)

				student, _ := MonitorManager.Students[offer.Subscribe]
				teacher, _ := MonitorManager.Teachers[offer.ID]

				teacher.sendHeaderData(student.headerData)
				student.dataDeal(msgDispatch.MsgData)
			}

		//msgFrom := message.MessageFrom{}
		//err := json.Unmarshal(msgFrom1, &msgFrom)
		//if err != nil{
		//	fmt.Println("man unmarshal error")
		//	continue
		//}
		//fmt.Println("man")
		//msgTo := message.MessageTo{
		//	From:msgFrom.From,
		//	Time : msgFrom.Time,
		//	Gid:"0",
		//	Content:msgFrom.Content,
		//}
		//
		//if msgFrom.Type == message.ClientMessage {
		//	if client, ok := manager.Clients[msgFrom.To]; ok {
		//		toMsg, _ := json.Marshal(msgTo)
		//
		//		client.Send <- toMsg
		//	} else {
		//		//redisSetMessage := message.RedisSetMessage{
		//		//	Key:message.Chant_Data,
		//		//	Message:msgFrom,
		//		//}
		//		//redisSetMessageByte, _ := json.Marshal(redisSetMessage)
		//		//var messageSlice = [][]byte{redisSetMessageByte}
		//		//err := rabbitmq.RedisSetProducter(messageSlice, message.Chant_Data)
		//		//if err != nil {
		//		//}
		//		WaitMsg[msgFrom.To] = append(WaitMsg[msgFrom.To],msgFrom.Content)
		//	}
		//}

		//if msgFrom.Type == message.GroupsMessage {
		//	groupClient, err := server.Redis.SMembers(msgFrom.To).Result()
		//	if err != nil {
		//		return
		//	}
		//
		//	msgTo.Gid = msgFrom.To
		//	msgToByte, _ := json.Marshal(msgTo)
		//
		//	if group, ok := manager.Groups[msgFrom.To]; ok {
		//		var messageSlice = [][]byte{}
		//		for _, clientId := range groupClient {
		//			if clientId == msgFrom.From {
		//				continue
		//			}
		//			if client, ok := group.Clients[clientId]; ok {
		//				client.Send <- msgToByte
		//			}else{
		//				msgFrom.To = clientId
		//				redisSetMessage := message.RedisSetMessage{
		//					Key:message.Chant_Data,
		//					Message:msgFrom,
		//				}
		//				redisSetMessageByte, _ := json.Marshal(redisSetMessage)
		//				messageSlice = append(messageSlice, redisSetMessageByte)
		//			}
		//		}
		//		err := rabbitmq.RedisSetProducter(messageSlice, message.Chant_Data)
		//		if err != nil {
		//		}
		//	}else{
		//		return
		//	}
		//}
		//case data := <-dispatch.Mideastream:
		//case data := <-dispatch.NetRateMonitor:

		}
	}
}


type monitorManager struct {
	Students map[int]*Student
	Teachers map[int]*Teacher
}

// Manager define a ws server manager
var MonitorManager = &monitorManager {
	Students: make(map[int]*Student, 5),
	Teachers: make(map[int]*Teacher, 100),
}

func createConnect(c *gin.Context) (*websocket.Conn, error) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
		ReadBufferSize:4000,
	}).Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func heart(){
	ticker := time.NewTicker(time.Second * 3)
	i := 0
	for{
		select {
		case <-ticker.C:
			for _, student := range MonitorManager.Students {
				student.calculateRate(i)
				student.Conn.WriteMessage(websocket.TextMessage, []byte("heart"))
			}
			for _, teacher := range MonitorManager.Teachers {
				teacher.calculateRate(i)
				teacher.Conn.WriteMessage(websocket.TextMessage, []byte("heart"))
			}
		}
	}
}

type rateManager struct {
	TeacherDownDataLen map[int]int
	TeacherUpDataLen map[int]int
	StudentDownDataLen map[int]int
	StudentUpDataLen map[int]int
}

var RateManager = &rateManager{
	TeacherDownDataLen : make(map[int]int, 5),
	TeacherUpDataLen : make(map[int]int, 5),
	StudentDownDataLen : make(map[int]int, 50),
	StudentUpDataLen : make(map[int]int, 50),
}
