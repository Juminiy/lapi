package server

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
	"time"
	"lapi/config"
	"lapi/middleware"
	"lapi/router"
	"lapi/storage"
)
const (
	CacheServerRedisPrefix = "SERVER_CACHE:"
)
type cacheServerRedis struct {
	Redis *redis.Client
}
func (redis *cacheServerRedis) Get(key string) ([]byte, error) {
	return redis.Redis.Get(CacheServerRedisPrefix+key).Bytes()
}

func (redis *cacheServerRedis) Set(key string, val []byte, ttl time.Duration) error {
	return redis.Redis.Set(CacheServerRedisPrefix+key,val,ttl).Err()
}

func (redis *cacheServerRedis) Delete(key string) error {
	return redis.Redis.Del(CacheServerRedisPrefix+key).Err()
}

func (redis *cacheServerRedis) Reset() error {
	prefixSet := redis.Redis.Scan(0,CacheServerRedisPrefix,1000)
	for _,key := range prefixSet.Args() {
		res := redis.Redis.Del(key.(string))
		if res.Err() != nil {
			return res.Err()
		}
	}
	return nil
}

func (redis *cacheServerRedis) Close() error {
	return redis.Redis.Close()
}

func CacheServer(app *fiber.App) {
	app.Use(cache.New(cache.Config{
		Next: 		  nil,
		Expiration:   30 * time.Second,
		CacheControl: true,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return utils.CopyString(ctx.Path())
		},
		Storage: 	  &cacheServerRedis{Redis: storage.Redis},
	}))
}

func InitServer() (app *fiber.App){
	app = fiber.New(config.ServerConfigs)
	middleware.Recover(app)
	middleware.CorsConfig(app)
	middleware.CsrfConfig(app)
	middleware.ConsoleLogger(app)
	router.RestApi(app)
	return app
}