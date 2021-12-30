package websocket

import (
	"github.com/gin-gonic/gin"
)

type AddGroupData struct {
	User_name string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email string `form:"email" binding:"required"`
	Code int `form:"code" binding:"required"`
}

// Client is a websocket client
type Group struct {
	ID     string
	Clients    map[string]*Client
	Send   chan []byte
}

func AddGroup(c *gin.Context)  {

}