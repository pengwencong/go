package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go/help"
	chatpb "go/proto/chat"
	"go/server"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"time"
	"unsafe"
)

const (
	// Address gRPC服务地址
	address = "127.0.0.1:50052"
)


var perSecondsReq int
var allReq int
var perReqMem = 1024 * 2

type Dispatch struct {
	Chat  	   chan []byte
	Mideastream    chan []byte
}


var Dispatcher = &Dispatch{
	Chat:  		make(chan []byte, 10),
	Mideastream:    make(chan []byte),
}

func (dispatch *Dispatch) Start() {
	//go heart()

	for {
		select {
		case chatData := <-dispatch.Chat:
			messageChat := &MessageChat{}
			err := json.Unmarshal(chatData, messageChat)
			if err != nil {

			}

			if client, ok := ChatManager.Clients[messageChat.To]; ok {
				messageSend := MessageSend{
					StringMessage,
					chatData,
				}

				client.Send <- messageSend
			} else if url, err := server.Redis.Get(string(messageChat.To)).Result(); err != nil {
				err := ChatServer.Send(url)
				if err != nil {

				}
			}

		}
	}
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

type chatManager struct {
	Clients map[int]*Client
	Rooms map[int]*Room
}

// Manager define a ws server manager
var ChatManager = &chatManager {
	Clients: make(map[int]*Client, 50),
	Rooms: make(map[int]*Room, 5),
}

type chatServer struct {

}

var ChatServer = chatServer{}

func (chatServer chatServer) Server() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		help.Log.Infof("chatserver Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	opts = append(opts,
		grpc.UnaryInterceptor(streamMonitor),
		)

	grpcServer := grpc.NewServer(opts...)
	grpc.EnableTracing = true

	chatpb.RegisterChatServer(grpcServer, ChatServer)

	fmt.Println("Listen on " + address)

	go calculateReq()

	err = grpcServer.Serve(listen)
	if err != nil {
		help.Log.Infof("chatserver server fail: %v", err)
	}
}

func (chatServer chatServer) ReciveData(ctx context.Context, req *chatpb.ChatRequest) (*chatpb.ChatResponse, error) {

	reqByte, err := json.Marshal(req)
	if err != nil {
		help.Log.Infof("chatserver server fail: %v", err)
		return nil, err
	}

	Dispatcher.Chat <- reqByte

	return &chatpb.ChatResponse{
		Status:200,
		Message:"ok",
	}, nil
}

func (chatServer chatServer) Send(url string) error {
	conn, err := grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		help.Log.Infof("chatserver send Dial: %v", err)
		return err
	}
	defer conn.Close()
	fmt.Println("afsd")
	fmt.Println(unsafe.Sizeof(conn))
	chatClient := chatpb.NewChatClient(conn)

	ctx := context.Background()

	stringData := &chatpb.StringData{
		From:1,
		To:2,
		Type:int64(USER),
		Datatype:int64(StringMessage),
		Data:"peng",
	}
	fileData := &chatpb.FileData{}
	chatRequest := &chatpb.ChatRequest{
		Data:stringData,
		File:fileData,
	}

	response, err := chatClient.ReciveData(ctx,
		chatRequest,
		grpc.WaitForReady(true),
	)
	if err != nil {
		help.Log.Infof("chatserver send function err: %v", err)
		return err
	}

	fmt.Println(response.Status)

	return nil
}

// interceptor 拦截器
func streamMonitor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	allReq++
	fmt.Println(*info)
	fmt.Println(req)
	fmt.Println(ctx)
	// 继续处理请求
	return handler(ctx, req)
}

func calculateReq(){
	tick := time.Tick(1 * time.Second)
	lastReq := 0
	for {
		select {
		case <- tick:
			if allReq != lastReq {
				perSecondsReq = allReq - lastReq
				lastReq = allReq
			}
		}
	}
}