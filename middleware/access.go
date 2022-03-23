package middleware

import (
	"github.com/gin-gonic/gin"
	"go/help"
)

func Access(c *gin.Context)  {
	help.Log.Info("access")

	c.Next()
}