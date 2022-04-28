package common

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fastrand"
	"strings"
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

func IsMobile(userAgent string) bool {
	if len(userAgent) == 0 {
		return false
	}

	isMobile := false
	mobileKeywords := []string{"Mobile", "Android", "Silk/", "Kindle", "BlackBerry", "Opera Mini", "Opera Mobi"}

	for _, word := range mobileKeywords {
		if strings.Contains(userAgent, word) {
			isMobile = true
			break
		}
	}

	return isMobile
}