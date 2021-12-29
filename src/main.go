package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go/router"
)

func main() {


	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)
	//禁用控制台颜色
	router.Init(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
