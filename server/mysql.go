package server

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go/help"
	"log"
	"sync"
)

type Mysql struct {
	GormDB *gorm.DB
}

var mysqlPool = sync.Pool{
	New :func() interface{} {
		return newMysql()
	},
}

var dns string

func InitMysqlConfig() {
	mysqlConfig := help.Conf.Mysql

	user := mysqlConfig.User
	password := mysqlConfig.Password
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	name := mysqlConfig.Name

	dns = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name)
}

func GetMysql() Mysql {
	return mysqlPool.Get().(Mysql)
}

func PutMysql(instance Mysql) {
	mysqlPool.Put(instance)
}

func InitMysqlPool(num int) error {
	for i:= 0; i < num; i++ {
		mysqlInstance := newMysql()
		if mysqlInstance.GormDB == nil {
			return errors.New("new mysql error")
		}
		mysqlPool.Put(mysqlInstance)
	}

	return nil
}

func newMysql() Mysql {

	MysqlSqlDB := Mysql{}

	var err error
	MysqlSqlDB.GormDB, err = gorm.Open("mysql", dns)
	if err != nil {
		log.Printf("open database error:%v", err)
		return MysqlSqlDB
	}

	MysqlSqlDB.GormDB.DB().SetMaxIdleConns(10)
	MysqlSqlDB.GormDB.DB().SetMaxOpenConns(100)
	MysqlSqlDB.GormDB.SingularTable(true)

	return MysqlSqlDB
}
