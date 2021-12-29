package router

import (
	"github.com/gin-gonic/gin"
	"go/controller"
)

func Init(engine *gin.Engine) {

	engine.GET("/video", controller.VideoStream)
}
