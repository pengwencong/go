package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
	"go/live"
	"go/monitor"
)

func Init(engine *gin.Engine) {

	live := engine.Group("/live")
	{
		live.GET("/userroom", controller.UserRoom)
		live.GET("/connectToRoom", live.ConnectToRoom)
		live.GET("/createRoom", controller.CreateRoom)
		live.GET("/roomInit", live.Init)
	}

	monitor := engine.Group("/monitor")
	{
		monitor.GET("/teacher", controller.Teacher)
		monitor.GET("/monitorStudent", monitor.MonitorStudent)
		monitor.GET("/student", controller.Student)
		monitor.GET("/studentConnect", monitor.StudentConnect)
	}

	engine.Static("/resource/video","./resource/video")
}
