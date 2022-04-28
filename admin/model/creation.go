package model

import "go/server"

type Technology struct {
	Name string `json:"name"`
	BookId int `json:"book_id"`
	OccupationId int `json:"occupation_id"`
	Desc string `json:"desc"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type TechnologyList struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BookName string `gorm:"-" json:"bookName"`
	OccupationName string `gorm:"-" json:"occupationName`
}

type Outline struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	BookId int `json:"book_id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type OutlineList struct {
	Title string `json:"title"`
}

type Plot struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	BookId int `json:"book_id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type PlotList struct {
	Title string `json:"title"`
}

type Design struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	BookId int `json:"book_id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type DesignList struct {
	Title string `json:"title"`
}

func GetDesignList(offset int, limit int) (designList []DesignList, err error) {
	mysql := server.GetMysql()

	err = mysql.GormDB.Table("design").
		Select("title").
		Where("id > ? ", offset).Limit(limit).Find(&designList).Error

	server.PutMysql(mysql)

	return designList, err
}

func (design Design) AddDesign() error {
	mysql := server.GetMysql()

	err := mysql.GormDB.Table("design").Create(design).Error

	server.PutMysql(mysql)

	if err != nil {
		return err
	}

	return nil
}

func GetPlotList(offset int, limit int) (plotList []PlotList, err error) {
	mysql := server.GetMysql()

	err = mysql.GormDB.Table("plot").
		Select("title").
		Where("id > ? ", offset).Limit(limit).Find(&plotList).Error

	server.PutMysql(mysql)

	return plotList, err
}

func (plot Plot) AddPlot() error {
	mysql := server.GetMysql()

	err := mysql.GormDB.Table("plot").Create(plot).Error

	server.PutMysql(mysql)

	if err != nil {
		return err
	}

	return nil
}

func GetOutlineList(offset int, limit int) (outlineList []OutlineList, err error) {
	mysql := server.GetMysql()

	err = mysql.GormDB.Table("outline").
		Select("title").
		Where("id > ? ", offset).Limit(limit).Find(&outlineList).Error

	server.PutMysql(mysql)

	return outlineList, err
}

func (outline Outline) AddOutline() error {
	mysql := server.GetMysql()

	err := mysql.GormDB.Table("outline").Create(outline).Error

	server.PutMysql(mysql)

	if err != nil {
		return err
	}

	return nil
}

func GetTechList(offset int, limit int) (techList []TechnologyList, err error) {
	mysql := server.GetMysql()

	err = mysql.GormDB.Table("technology t").
		Select("t.id, t.name, b.name AS book_name, o.name AS occupation_name").
		Joins("left join book b on b.id = t.book_id").
		Joins("left join occupation o on o.id = t.occupation_id").
		Where("t.id > ? ", offset).Limit(limit).Find(&techList).Error

	server.PutMysql(mysql)

	return techList, err
}

func SearchTech(name string, occupationId int) (techList []TechnologyList, err error) {
	mysql := server.GetMysql()

	err = mysql.GormDB.Table("technology").
		Select("name").
		Where("name like ?", name + "%").
		Where("occupation_id = ?", occupationId).
		Find(&techList).Error

	server.PutMysql(mysql)

	return techList, err
}

func (tech Technology) AddTech() error {
	mysql := server.GetMysql()

	err := mysql.GormDB.Table("technology").Create(tech).Error

	server.PutMysql(mysql)

	if err != nil {
		return err
	}

	return nil
}
