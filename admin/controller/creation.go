package controller

import (
	"github.com/gin-gonic/gin"
	"go/admin/model"
	"go/common"
	"go/help"
	"strconv"
	"time"
)

func AddDesign(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.PostForm("bookId"))
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	now := time.Now().Unix()
	design := model.Design{
		Title: title,
		BookId: bookId,
		Desc: desc,
		CreateTime: now,
		UpdateTime: now,
	}

	err := design.AddDesign()
	if err != nil {
		help.Log.Infof("bookId:%s, name:%s, desc:%s create design fail: %s", bookId,
			title, desc, err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	common.Json(c, 200, H)
}

func DesignList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.PostForm("from"))
	size := 10

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	designList, err := model.GetDesignList(offset, size)

	if err != nil {
		help.Log.Infof("get techList fail: %s", err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	H["data"] = designList
	common.Json(c, 200, H)
}

func AddPlot(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.PostForm("bookId"))
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	now := time.Now().Unix()
	plot := model.Plot{
		Title: title,
		BookId: bookId,
		Desc: desc,
		CreateTime: now,
		UpdateTime: now,
	}

	err := plot.AddPlot()
	if err != nil {
		help.Log.Infof("bookId:%s, name:%s, desc:%s create plot fail: %s", bookId,
			title, desc, err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	common.Json(c, 200, H)
}

func PlotList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.PostForm("from"))
	size := 10

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	outlineList, err := model.GetPlotList(offset, size)

	if err != nil {
		help.Log.Infof("get techList fail: %s", err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	H["data"] = outlineList
	common.Json(c, 200, H)
}

func AddOutline(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.PostForm("bookId"))
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	now := time.Now().Unix()
	tech := model.Outline{
		Title: title,
		BookId: bookId,
		Desc: desc,
		CreateTime: now,
		UpdateTime: now,
	}

	err := tech.AddOutline()
	if err != nil {
		help.Log.Infof("bookId:%s, name:%s, desc:%s create outline fail: %s", bookId,
			title, desc, err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	common.Json(c, 200, H)
}

func OutlineList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.PostForm("from"))
	size := 10

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	outlineList, err := model.GetOutlineList(offset, size)

	if err != nil {
		help.Log.Infof("get techList fail: %s", err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	H["data"] = outlineList
	common.Json(c, 200, H)
}

func AddTech(c *gin.Context) {
	occupationId, _ := strconv.Atoi(c.PostForm("occupationId"))
	bookId, _ := strconv.Atoi(c.PostForm("bookId"))
	techName := c.PostForm("techName")
	techDesc := c.PostForm("techDesc")

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	now := time.Now().Unix()
	tech := model.Technology{
		Name: techName,
		BookId: bookId,
		OccupationId: occupationId,
		Desc: techDesc,
		CreateTime: now,
		UpdateTime: now,
	}

	err := tech.AddTech()
	if err != nil {
		help.Log.Infof("bookId:%s, occupationId:%s, name:%s, desc:%s create tech fail: %s",
			bookId, occupationId, techName, techDesc, err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	common.Json(c, 200, H)

}

func TechList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.PostForm("from"))
	size := 10

	H := gin.H{
		"status":200,
		"message":"ok",
	}

	techList, err := model.GetTechList(offset, size)

	if err != nil {
		help.Log.Infof("get techList fail: %s", err.Error())

		H["status"] = 500
		H["message"] = err.Error()
		common.Json(c, 200, H)

		return
	}

	H["data"] = techList
	common.Json(c, 200, H)
}


