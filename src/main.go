package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go/router"
	"go/server"
)

func main() {

	err := server.InitRedis()
	if err != nil {
		fmt.Println(err)
		return
	}

	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)
	//禁用控制台颜色
	router.Init(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
