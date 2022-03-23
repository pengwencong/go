package middleware

import (
	"github.com/gin-gonic/gin"
	"go/common"
	"go/help"
	"go/server"
)

func Login(c *gin.Context)  {
	token := c.PostForm("token")

	redis := server.GetRedis()
	_, err := redis.Get(token).Result()
	server.PutRedis(redis)
	if err != nil {
		help.Log.Infof("redis get token error: %s", err.Error())
		common.Json(c, 200, gin.H{
			"status":300,
			"message":"no login",
		})
	}

	c.Next()
}