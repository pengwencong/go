package router

import (
	"github.com/gin-gonic/gin"
	adcontroller "go/admin/controller"
	"go/admin/middleware"
	mw "go/middleware"
	"go/chat"
	"go/controller"
	"go/elastic"
	"go/live"
	"go/monitor"
)

func Init(engine *gin.Engine) {

	adminGroup := engine.Group("/admin")
	{
		adminGroup.POST("/login", adcontroller.Login)
		adminGroup.POST("/isLogin", adcontroller.IsLogin)
		adminGroup.POST("/creation", middleware.Login, adcontroller.Creation)
	}

	elasticGroup := engine.Group("/elastic").Use(mw.Access)
	{
		elasticGroup.GET("/create", elastic.CreateIndex)
		elasticGroup.GET("/test", elastic.TestSearch)
	}

	liveGroup := engine.Group("/live").Use(mw.Access)
	{
		liveGroup.GET("/userroom", controller.UserRoom)
		liveGroup.GET("/connectToRoom", live.ConnectToRoom)
		liveGroup.GET("/createRoom", controller.CreateRoom)
		liveGroup.GET("/roomInit", live.Init)
	}

	monitorGroup := engine.Group("/monitor").Use(mw.Access)
	{
		monitorGroup.GET("/teacher", controller.Teacher)
		monitorGroup.GET("/monitorStudent", monitor.MonitorStudent)
		monitorGroup.GET("/student", controller.Student)
		monitorGroup.GET("/studentConnect", monitor.StudentConnect)
	}

	chatGroup := engine.Group("/chat").Use(mw.Access)
	{
		chatGroup.GET("/login", controller.Login)
		chatGroup.GET("/connect", chat.Connect)
	}

	engine.Static("/resource/video","./resource/video")
	engine.Static("/public/static","./public/static")
	engine.StaticFile("/admin","./views/admin/index.html")
	engine.StaticFile("/shopify","./views/shopify.html")
	engine.Static("/resource/js","./resource/js")
	engine.GET("/monitorGc", controller.MonitorGC)
	engine.GET("/test", controller.TestSlice)
}
