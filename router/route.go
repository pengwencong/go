package router

import (
	"github.com/gin-gonic/gin"
	"go/chat"
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

	chatGroup := engine.Group("/chat")
	{
		chatGroup.GET("/login", controller.Login)
		chatGroup.GET("/connect", chat.Connect)
	}

	engine.Static("/resource/video","./resource/video")
	engine.Static("/resource/css","./resource/css")
	engine.Static("/resource/js","./resource/js")
	engine.GET("/monitorGc", controller.MonitorGC)
}
