package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/help"
	"go/router"
	"go/server"
)

func main() {
	help.InitZap()

	err := help.InitYaml()
	if err != nil {
		help.Log.Error("init yaml err: ", err.Error())
		return
	}

	server.InitMysqlConfig()

	err = server.InitMysqlPool(2)
	if err != nil {
		help.Log.Error("init mysql err: ", err.Error())
		return
	}

	server.InitRedisConfig()

	err = server.InitRedisPool(2)
	if err != nil {
		help.Log.Error("init redis err: ", err.Error())
		return
	}
	//
	//server.InitEsConfig()
	//
	//err = server.InitEsPool(2)
	//if err != nil {
	//	help.Log.Error("init es err: ", err.Error())
	//	return
	//}
	//
	//server.InitKafkaConfig()
	//err = server.InitKafkaPool(2)
	//if err != nil {
	//	help.Log.Error("init kafka err: ", err.Error())
	//	return
	//}

	//go server.ConsumerGroup()
	//go server.Consumer()
	//go live.Dispatcher.Start()
	//go monitor.Dispatcher.Start()
	//go chat.Dispatcher.Start()
	//go chat.ChatServer.Server()
	//chat.ChatManager.Clients[0] = nil
	//
	//server.GCTicker()

	engine := gin.Default()
	//pprof.Register(engine)

	router.Init(engine)

	err = engine.Run(":8080")
	//err = engine.RunTLS(":443", "./runtime/tls/server.pem", "./runtime/tls/server.key")
	fmt.Println("listen err:", err)
}

