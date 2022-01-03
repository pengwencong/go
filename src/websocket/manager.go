package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/message"
	"log"
	"net/http"
	"os"
	"time"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[string]*Client
	Groups	   map[string]*Group
	Chat  	   chan []byte
	Monitor    chan []byte
	Register   chan *Client
	Unregister chan *Client
}


// Manager define a ws server manager
var Manager = &ClientManager{
	Chat:  		make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
	Groups:    	make(map[string]*Group),
	Monitor:    make(chan []byte),
}

var WaitMsg = make(map[string][]string)

func Connect(c *gin.Context){
	conn := CreateConnect(c)
	if conn == nil {
		return
	}
	conn.SetCloseHandler(CloseHandle)
	ss := websocket.Subprotocols(c.Request)
	for _, va := range ss {
		if va == ""{

		}
	}
	_, messag, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return
	}
	if _, ok := Manager.Clients[string(messag)]; ok {
		return
	}
	//group := &Group{
	//	ID:"g1",
	//	Clients: make(map[string]*Client),
	//	Send: make(chan []byte),
	//}


	// websocket connect
	client := &Client{
		ID: string(messag),
		Socket: conn,
		Send: make(chan []byte, 20),
		HeartTime: 0,
		DataBuff : []byte{},
	}

	go client.Read()
	go client.Write()

	//group.Clients[string(message)] = client



	manager := GetManager()
	//manager.Groups["g1"] = group
	manager.Register <- client


	if val, ok := WaitMsg[client.ID]; ok {
		fmt.Println(client.ID, val)
		for _, v := range val {
			from := "1"
			if client.ID == "1" {
				from = "2"
			}
			msgTo := message.MessageTo{
				From:from,
				Time : "11:20",
				Gid:"0",
				Content:v,
			}
			toMsg, _ := json.Marshal(msgTo)

			client.Send <- toMsg
		}
	}
	delete(WaitMsg,client.ID)
}



//func sendBackLogMsg(client *Client) error{
//	for {
//		msgFrom := message.MessageFrom{}
//		msgTo := message.MessageTo{}
//
//		err = json.Unmarshal([]byte(data), &msgFrom)
//		if err != nil {
//			return err
//		}
//
//		msgTo.From = msgFrom.From
//		msgTo.Time = msgFrom.Time
//		msgTo.Gid  = "0"
//		msgTo.Content = msgFrom.Content
//		msgToByte, _ := json.Marshal(msgTo)
//
//		client.Send <- msgToByte
//	}
//}

func GetManager() *ClientManager{
	return Manager
}

func Diala(c *gin.Context){
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	err = conn.WriteMessage(websocket.PingMessage, []byte{})
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func CreateConnect(c *gin.Context) *websocket.Conn {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
		ReadBufferSize:4000,
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil
	}

	return conn
}

func CloseHandle(code int, text string) error{
	return nil
}

// Start is to start a ws server
func (manager *ClientManager) Start() {
	//ticker := time.NewTicker(time.Second * 30)
	//go heart(ticker, Manager)

	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn.ID] = conn
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn.ID]; ok {
				close(conn.Send)
				conn.Socket.Close()
				delete(manager.Clients, conn.ID)
			}
		case msgFrom1 := <-manager.Chat:
			if client, ok := manager.Clients["1"]; ok {
				client.Socket.WriteMessage(websocket.BinaryMessage, msgFrom1)
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
		case binaryData := <-manager.Monitor:

			f, _ := os.OpenFile("2.mp3", os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer f.Close()
			f.Write(binaryData)
		}
	}
}

func heart(ticker *time.Ticker, manager *ClientManager){
	for{
		select {
		case <-ticker.C:
			for _, client := range manager.Clients{
				if(client.HeartTime >= 2){
					Manager.Unregister <- client
					continue
				}
				client.Socket.WriteJSON(message.MessageTo{From: "0", Time:"0", Gid:"0", Content: "ping"})
				client.HeartTime += 1
			}
		}
	}
}
