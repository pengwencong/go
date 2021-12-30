package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go/router"
	"go/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	hs := &http.Server{
		Addr:":80",
		Handler: engine,
	}

	go hs.ListenAndServeTLS("./runtime/tls/server.pem", "./runtime/tls/server.key")
	//engine.RunTLS(":80", )

	restart(hs)
}

func restart(hs *http.Server){
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <- sigs:
		hs.Shutdown(context.TODO())
	}
}
