package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
	"go/websocket"
)

func Init(engine *gin.Engine) {

	engine.GET("/video", controller.VideoStream)
	engine.GET("/video1", controller.VideoStream1)
	engine.GET("/redis", controller.RedisOperate)
	engine.GET("/ws", websocket.Connect)
	engine.StaticFile("/1.mp4","./resource/video/1.mp4")
}
