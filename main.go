package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go/help"
	"go/live"
	"go/router"
	"go/server"
)

func main() {
	help.InitZap()

	err := server.InitRedis()
	if err != nil {
		help.Log.Info("init redis err:", err.Error())
		return
	}

	go live.Dispatcher.Start()

	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)

	router.Init(engine)

	err = engine.RunTLS(":443", "./runtime/tls/server.pem", "./runtime/tls/server.key")
	fmt.Println("listen err:", err)
}

