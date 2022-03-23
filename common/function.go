package common

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fastrand"
)

func GenRandStr(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := fastrand.Uint32n(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func Json(c *gin.Context, status int, h gin.H) {
	c.JSON(status,h)
}