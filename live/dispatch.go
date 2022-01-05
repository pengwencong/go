package live

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/message"
	"net/http"
)

// ClientManager is a websocket manager
type Dispatch struct {
	Chat  	   chan message.MessageDispatch
	Mideastream    chan []byte
	Action chan []byte
}


var Dispatcher = &Dispatch{
	Chat:  		make(chan message.MessageDispatch),
	Mideastream:    make(chan []byte),
	Action: make(chan []byte),
}


func (dispatch *Dispatch) Start() {
	//ticker := time.NewTicker(time.Second * 30)
	//go heart(ticker, Manager)

	for {
		select {
		case msgDispatch := <-dispatch.Chat:
			switch msgDispatch.Type {
			case message.OfferMessage:
				offer := message.MessageOffer{}
				json.Unmarshal(msgDispatch.MsgSend.Data, &offer)

				room, _ := LiveManager.Rooms[offer.Subscribe]
				client, _ := LiveManager.Clients[offer.ID]
				fmt.Printf("%+v\n", room)
				fmt.Printf("%+v\n", client)
				client.sendHeaderData(room.headerData)
				room.dataDeal(msgDispatch.MsgSend)
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

		}
	}
}


type liveManager struct {
	Rooms map[int]*Room
	Clients map[int]*Client
}

// Manager define a ws server manager
var LiveManager = &liveManager{
	Rooms: make(map[int]*Room, 5),
	Clients: make(map[int]*Client, 100),
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


//func heart(ticker *time.Ticker, manager *ClientManager){
//	for{
//		select {
//		case <-ticker.C:
//			for _, client := range manager.Clients{
//				if(client.HeartTime >= 2){
//					continue
//				}
//				//client.Socket.WriteJSON(message.MessageTo{From: "0", Time:"0", Gid:"0", Content: "ping"})
//				//client.HeartTime += 1
//			}
//		}
//	}
//}

