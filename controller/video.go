package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/help"
	"go/live"
	"runtime"
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


	i := 1
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <- tick:
			var memStats = &runtime.MemStats{}
			runtime.ReadMemStats(memStats)

			fmt.Printf("BEFORE GC: %+v\n", memStats)

			var GCstats = &debug.GCStats{}
			debug.ReadGCStats(GCstats)

			fmt.Printf("%+v\n", GCstats)
			i++
			if i < 4 {
				var test = make([]byte, 1024*1024 * 2)
				if len(test) == 20 {}
			} else {
				var test = make([]byte, 1024*1024)
				if len(test) == 20 {}
			}

		}
	}
}
