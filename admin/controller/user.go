package controller

import (
	"github.com/gin-gonic/gin"
	"go/admin/model"
	"go/common"
	"go/help"
	"go/server"
	"time"
)

var tokenExpire time.Duration = time.Hour * 24 * 7

func IsLogin(c *gin.Context) {
	token := c.PostForm("token")

	redis := server.GetRedis()
	_, err := redis.Get(token).Result()
	server.PutRedis(redis)
	H := gin.H{
		"status":200,
		"message":"ok",
	}

	if err != nil {
		help.Log.Infof("redis get token error: %s", err.Error())
		common.Json(c, 200, gin.H{
			"status":300,
			"message":"no login",
		})
		return
	}

	common.Json(c, 200, H)
}

func Login(c *gin.Context){
	phione := c.PostForm("phione")
	password := c.PostForm("password")

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	admin := model.Admin{}
	err := admin.GetAdmin(phione)
	if err != nil {
		help.Log.Infof("phione:%s, password:%s login fail: %s", phione, password, err.Error())

		H["status"] = 300
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	err = admin.IsAdmin(password)
	if err != nil {
		help.Log.Infof("phione:%s, password:%s login fail: %s", phione, password, err.Error())

		H["status"] = 300
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	jwt := help.NewJWT(string(admin.Id) + admin.PrivateKey)
	token, err := jwt.CreateToken()
	if err != nil {
		help.Log.Infof("phione:%s, password:%s login fail: %s", phione, password, err.Error())

		H["status"] = 300
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	redis := server.GetRedis()
	_, err = redis.Set(token, 1, tokenExpire).Result()
	server.PutRedis(redis)
	if err != nil {
		help.Log.Infof("phione:%s, password:%s login fail: %s", phione, password, err.Error())

		H["status"] = 300
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	H["token"] = token
	common.Json(c, 200, H)
}

func Index(c *gin.Context){
	isMobile := common.IsMobile(c.GetHeader("User-Agent"))

	if isMobile {
		c.HTML(200,"indexMobile.html",gin.H{})
	}else{
		c.HTML(200,"index.html",gin.H{})
	}

}