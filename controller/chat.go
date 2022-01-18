package controller

import (
	"github.com/gin-gonic/gin"
	"go/help"
	"strconv"
)

func Login(c *gin.Context) {
	ID, err := strconv.Atoi( c.Query("id") )
	if err != nil {
		help.Log.Info("create teacher Atoi err:", err.Error())
	}
	name := "peng"
	chatName := "wen"

	c.HTML(200,"chat.html",gin.H{
		"ID": ID,
		"name": name,
		"chatName":chatName,
	})
}