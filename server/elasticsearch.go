package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go/help"
	"sync"
)

var user string
var password string
var addr string

var ClientES *elastic.Client
var ctx = context.Background()

var EsPool = sync.Pool{
	New :func() interface{} {
		return newES()
	},
}

func newES() *elastic.Client {
	ClientES, err := elastic.NewClient(
		elastic.SetURL(addr),
		elastic.SetBasicAuth(user,password),
		elastic.SetSniff(false),
	)
	if err != nil {
		help.Log.Infof("es new client error: %s", err.Error())
		return nil
	}

	return ClientES
}

func InitEsPool(num int) error {
	for i:= 0; i < num; i++ {
		esInstance := newES()
		if esInstance == nil {
			return errors.New("new redis error")
		}
		EsPool.Put(esInstance)
	}

	return nil
}

func InitEsConfig() {
	esConfig := help.Conf.Es

	host := esConfig.Host
	port := esConfig.Port
	user = esConfig.User
	password = esConfig.Password

	addr = fmt.Sprintf("http://%s:%s", host, port)
}



func GetEs() *elastic.Client {
	return EsPool.Get().(*elastic.Client)
}

func PutEs(instance *elastic.Client) {
	EsPool.Put(instance)
}