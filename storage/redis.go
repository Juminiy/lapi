package storage

import (
	"github.com/go-redis/redis"
	"strconv"
	"lapi/config"
)

var (
	Redis *redis.Client
)

func RedisConnect() error {
	db,_ := strconv.ParseInt(config.Config("REDIS_DB"),10,64)
	client := redis.NewClient(&redis.Options{
		Addr: config.Config("REDIS_ADDR"),
		Password: config.Config("REDIS_PASS"),
		DB: int(db),
	})
	_,err := client.Ping().Result()
	if err != nil {
		return nil
	}
	Redis = client
	return nil
}

