package controller

import (
	"github.com/gin-gonic/gin"
)

func VideoStream(c *gin.Context){
	c.HTML(200,"video.html",gin.H{})
}