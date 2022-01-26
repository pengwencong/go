package server

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() error {

	addr := fmt.Sprintf("%s:%s", "120.77.17.51", "6379")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Redis.Ping().Result()
	return err
}

func Set(key string, value string, ) error {


	return nil
}
