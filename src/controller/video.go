package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/server"
)

func VideoStream(c *gin.Context){
	id := c.Query("id")
	//c.JSON(200,gin.H{
	//	"message":"success",
	//})
	c.HTML(200,"video1.html",gin.H{
		"id":id,
	})
}

func RedisOperate(c *gin.Context){
	result, err := server.Redis.Set("peng","wenport",0).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	result, err = server.Redis.Get("peng").Result()
	fmt.Println(result)
}