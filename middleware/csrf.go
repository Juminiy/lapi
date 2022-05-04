package middleware

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"time"
	envConfig "zhaoxin-api/config"
	"zhaoxin-api/storage"
	envUtils "zhaoxin-api/utils"
)
const (
	CsrfStoreRedisPrefix = "CSRF_STORE:"
)
type CsrfStorage struct {
	Redis *redis.Client
}
func (redis *CsrfStorage) Get(key string) ([]byte, error) {
	return redis.Redis.Get(CsrfStoreRedisPrefix+key).Bytes()
}

func (redis *CsrfStorage) Set(key string, val []byte, ttl time.Duration) error {
	return redis.Redis.Set(CsrfStoreRedisPrefix+key,val,ttl).Err()
}

func (redis *CsrfStorage) Delete(key string) error {
	return redis.Redis.Del(CsrfStoreRedisPrefix+key).Err()
}

func (redis *CsrfStorage) Reset() error {
	prefixSet := redis.Redis.Scan(0,CsrfStoreRedisPrefix,1000)
	for _,key := range prefixSet.Args() {
		res := redis.Redis.Del(key.(string))
		if res.Err() != nil {
			return res.Err()
		}
	}
	return nil
}
func (redis *CsrfStorage) Close() error {
	return redis.Redis.Close()
}

func CsrfConfig(app *fiber.App) {
	app.Use(csrf.New(csrf.Config{
		Next: nil,
		KeyLookup: "header:X-Csrf-Token",
		CookieName: "x_csrf_token",
		CookieDomain: envConfig.Config("BACKEND_COOKIE_DOMAIN"),
		CookiePath: "/",
		//CookieSecure: true,
		//CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		Expiration: 1 * time.Hour,
		Storage: &CsrfStorage{Redis: storage.Redis},
		KeyGenerator: utils.UUID,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return envUtils.RequestFailureResponse(ctx,"CSRF Protection Rejected")
		},
	}))
}