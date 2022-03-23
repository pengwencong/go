package server

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go/help"
	"sync"
)

var redisOption *redis.Options

var redisPool = sync.Pool{
	New :func() interface{} {
		return newRedis()
	},
}

func GetRedis() *redis.Client {
	return redisPool.Get().(*redis.Client)
}

func PutRedis(instance *redis.Client) {
	redisPool.Put(instance)
}

func InitRedisConfig() {
	redisConfig := help.Conf.Redis

	password := redisConfig.Password
	host := redisConfig.Host
	port := redisConfig.Port
	db := redisConfig.DB

	addr := fmt.Sprintf("%s:%s", host, port)
	redisOption = &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
}

func InitRedisPool(num int) error {
	for i:= 0; i < num; i++ {
		redisInstance := newRedis()
		if redisInstance == nil {
			return errors.New("new redis error")
		}
		redisPool.Put(redisInstance)
	}

	return nil
}

func newRedis() *redis.Client {
	redisinstance := redis.NewClient(redisOption)

	_, err := redisinstance.Ping().Result()

	if err != nil {
		help.Log.Infof("new redis ping error :%s", err.Error())
		redisinstance = nil
	}
	return redisinstance
}


func Set(key string, value string, ) error {


	return nil
}
