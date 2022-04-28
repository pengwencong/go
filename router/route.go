package router

import (
	"github.com/gin-gonic/gin"
	adcontroller "go/admin/controller"
	"go/admin/middleware"
	"go/chat"
	"go/controller"
	"go/elastic"
	"go/live"
	mw "go/middleware"
	"go/monitor"
	"go/server"
)

func Init(engine *gin.Engine) {
	adminGroup := engine.Group("/admin")
	{
		adminGroup.GET("/index", adcontroller.Index)
		adminGroup.POST("/login", adcontroller.Login)
		adminGroup.POST("/isLogin", adcontroller.IsLogin)
		creationGroup := adminGroup.Group("/creation").Use(middleware.Login)
		{
			creationGroup.POST("/addTech", adcontroller.AddTech)
			creationGroup.POST("/addTechScene", adcontroller.AddTechScene)
			creationGroup.POST("/techList", adcontroller.TechList)
			creationGroup.POST("/searchTech", adcontroller.SearchTech)
			creationGroup.POST("/outlineList", adcontroller.OutlineList)
			creationGroup.POST("/addOutline", adcontroller.AddOutline)
			creationGroup.POST("/plotList", adcontroller.PlotList)
			creationGroup.POST("/addPlot", adcontroller.AddPlot)
			creationGroup.POST("/designList", adcontroller.DesignList)
			creationGroup.POST("/addDesign", adcontroller.AddDesign)
		}
	}

	kafkaGroup := engine.Group("/kafka")
	{
		kafkaGroup.GET("/producter", server.Producer)
		kafkaGroup.GET("/test/*name", controller.RedisTest)
	}

	elasticGroup := engine.Group("/elastic").Use(mw.Access)
	{
		elasticGroup.GET("/create", elastic.CreateIndex)
		elasticGroup.GET("/test", elastic.TestSearch)
		elasticGroup.GET("/ag",elastic.Aggregation)
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
	//engine.StaticFile("/admin","./views/admin/index.html")
	engine.Static("/resource/js","./resource/js")
	engine.GET("/monitorGc", controller.MonitorGC)
	engine.GET("/test", controller.TestSlice)

}
