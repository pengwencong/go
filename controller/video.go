package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/help"
	"go/live"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"
)

func UserRoom(c *gin.Context) {
	roomID, _ := strconv.Atoi( c.Query("roomID") )
	userID := 1

	if _, ok := live.LiveManager.Rooms[roomID]; !ok {
	}

	c.HTML(200,"userroom.html",gin.H{
		"roomID": roomID,
		"userID": userID,
	})
}

func CreateRoom(c *gin.Context) {
	roomID, err := strconv.Atoi( c.Query("roomID") )
	if err != nil {
		help.Log.Info("create room Atoi err:", err.Error())
	}

	c.HTML(200,"liveroom.html",gin.H{
		"roomID": roomID,
	})
}

func Student(c *gin.Context) {
	studentID, err := strconv.Atoi( c.Query("studentID") )
	if err != nil {
		help.Log.Info("create student Atoi err:", err.Error())
	}

	c.HTML(200,"student.html",gin.H{
		"studentID": studentID,
	})
}

func Teacher(c *gin.Context) {
	teacherID, err := strconv.Atoi( c.Query("teacherID") )
	if err != nil {
		help.Log.Info("create teacher Atoi err:", err.Error())
	}
	studentID := 1

	c.HTML(200,"teacher.html",gin.H{
		"teacherID": teacherID,
		"studentID": studentID,
	})
}

func MonitorGC(c *gin.Context) {
	tick := time.Tick(3 * time.Minute)

	for {
		select {
		case <- tick:
			var gcStatus = &debug.GCStats{}
			debug.ReadGCStats(gcStatus)

			fmt.Println(gcStatus)
		}
	}
}

func TestSlice(c *gin.Context) {
	url := "https://www.dddbiquge.cc/chapter/45082839_7130174.html"

	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Acceept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")

	res, err := client.Do(request)
	if err != nil {

	}
	defer res.Body.Close()

	if res.StatusCode != 200 {

	}

	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

