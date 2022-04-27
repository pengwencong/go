package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RedisTest(c *gin.Context) {

}

func getVal(i int) int {
	fmt.Println(i)
	return i
}