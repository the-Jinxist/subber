package config

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

var rdb *redis.Pool

func InitRedis() {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	rdb = redisPool
}

func GetRedis() *redis.Pool {
	return rdb
}
