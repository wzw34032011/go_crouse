package dao

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisConf struct {
	network  string
	host     string
	port     int
	password string
	db       int
}

func initRedis(c *RedisConf) {
	options := redis.Options{
		Network:  c.network,
		Addr:     fmt.Sprintf("%s:%d", c.host, c.port),
		Password: c.password,
		DB:       c.db,
	}

	ClientRedis = redis.NewClient(&options)
}
