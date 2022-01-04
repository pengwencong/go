package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
	"go/help"
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

	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)

	router.Init(engine)

	err = engine.RunTLS(":443", "./runtime/tls/server.pem", "./runtime/tls/server.key")
	fmt.Println("listen err:", err)
}

