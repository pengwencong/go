package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
	"go/websocket"
)

func Init(engine *gin.Engine) {

	engine.GET("/video", controller.VideoStream)
	engine.GET("/redis", controller.RedisOperate)
	engine.GET("/ws", websocket.Connect)
}
