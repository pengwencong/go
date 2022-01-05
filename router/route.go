package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
	"go/live"
	"go/monitor"
)

func Init(engine *gin.Engine) {

	engine.Group("/live")
	{
		engine.GET("/userroom", controller.UserRoom)
		engine.GET("/connectToRoom", live.ConnectToRoom)
		engine.GET("/createRoom", controller.CreateRoom)
		engine.GET("/roomInit", live.Init)
	}

	engine.Group("/monitor")
	{
		engine.GET("/teacher", controller.Teacher)
		engine.GET("/monitorStudent", monitor.MonitorStudent)
		engine.GET("/student", controller.Student)
		engine.GET("/studentConnect", monitor.StudentConnect)
	}

	engine.Static("/resource/video","./resource/video")
}
