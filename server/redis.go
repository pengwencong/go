package server

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go/help"
	"sync"
)

var redisOption *redis.ClusterOptions

var redisPool = sync.Pool{
	New :func() interface{} {
		return newRedis()
	},
}

func GetRedis() *redis.ClusterClient {
	return redisPool.Get().(*redis.ClusterClient)
}

func PutRedis(instance *redis.ClusterClient) {
	redisPool.Put(instance)
}

func InitRedisConfig() {
	redisConfig := help.Conf.Redis

	password := redisConfig.Password
	host := redisConfig.Host
	port := redisConfig.Port

	addr := fmt.Sprintf("%s:%s", host, port)
	addr1 := fmt.Sprintf("%s:%s", "106.55.178.129", "6380")
	addr2 := fmt.Sprintf("%s:%s", "120.77.17.51", "6379")
	redisOption = &redis.ClusterOptions {
		Addrs:    []string{addr, addr1, addr2},
		PoolSize:3,
		Password: password,
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

func newRedis() *redis.ClusterClient {
	redisinstance := redis.NewClusterClient(redisOption)

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
