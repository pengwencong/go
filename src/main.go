package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
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
	engine.Use(TlsHandler())
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)
	//禁用控制台颜色
	router.Init(engine)

	hs := &http.Server{
		Addr:":80",
		Handler: engine,
	}

	err = hs.ListenAndServeTLS("./runtime/tls/server.pem", "./runtime/tls/server.key")
	//engine.RunTLS(":80", )

	fmt.Println("listen err:", err)

	go restart(hs)
}

func restart(hs *http.Server){
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <- sigs:
		hs.Shutdown(context.TODO())
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":80",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
