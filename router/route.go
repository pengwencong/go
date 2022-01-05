package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
	"go/live"
	"go/monitor"
)

func Init(engine *gin.Engine) {

	liveGroup := engine.Group("/live")
	{
		liveGroup.GET("/userroom", controller.UserRoom)
		liveGroup.GET("/connectToRoom", live.ConnectToRoom)
		liveGroup.GET("/createRoom", controller.CreateRoom)
		liveGroup.GET("/roomInit", live.Init)
	}

	monitorGroup := engine.Group("/monitor")
	{
		monitorGroup.GET("/teacher", controller.Teacher)
		monitorGroup.GET("/monitorStudent", monitor.MonitorStudent)
		monitorGroup.GET("/student", controller.Student)
		monitorGroup.GET("/studentConnect", monitor.StudentConnect)
	}

	engine.Static("/resource/video","./resource/video")
}
