package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/help"
	"go/server"
	"time"
)

func RedisTest(c *gin.Context) {
	redis := server.GetRedis()
	res, err := redis.Set("peng", "wen",-1).Result()
	if err != nil {
		help.Log.Errorf("redis set error: %s\n", err.Error())
	}

	fmt.Println(res)

	res, err = redis.Set("peng", "wen", time.Second * 60 * 10).Result()
	if err != nil {
		help.Log.Errorf("redis set error: %s\n", err.Error())
	}

	fmt.Println(res)

	res, err = redis.Get("peng").Result()
	if err != nil {
		help.Log.Errorf("redis get error: %s\n", err.Error())
	}

	fmt.Println(res)
}