package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go/router"
	"go/server"
	"go/websocket"
)

func main() {

	err := server.InitRedis()
	if err != nil {
		fmt.Println(err)
		return
	}

	go websocket.Manager.Start()
	
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)
	//禁用控制台颜色
	router.Init(engine)

	err = engine.RunTLS(":443", "./runtime/tls/server.pem", "./runtime/tls/server.key")

	fmt.Println("listen err:", err)
}

