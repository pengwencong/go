package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
	"go/live"
)

func Init(engine *gin.Engine) {

	engine.GET("/userroom", controller.UserRoom)
	engine.GET("/roomInit", live.Init)
	engine.GET("/connectToRoom", live.ConnectToRoom)
	engine.GET("/createRoom", controller.CreateRoom)
	engine.Static("/resource/video","./resource/video")
}
