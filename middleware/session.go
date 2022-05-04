package middleware

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"time"
	envConfig "lapi/config"
	"lapi/storage"
)

const (
	SessionStoreRedisPrefix = "SESSION_STORE:"
)
var (
	MySessionStore *session.Store
)

type sessionStorage struct {
	Redis *redis.Client
}
func init() {
	SessionStorage()
}
func (redis *sessionStorage) Get(key string) ([]byte, error) {
	return redis.Redis.Get(SessionStoreRedisPrefix+key).Bytes()
}

func (redis *sessionStorage) Set(key string, val []byte, ttl time.Duration) error {
	return redis.Redis.Set(SessionStoreRedisPrefix+key,val,ttl).Err()
}

func (redis *sessionStorage) Delete(key string) error {
	return redis.Redis.Del(SessionStoreRedisPrefix+key).Err()
}

func (redis *sessionStorage) Reset() error {
	prefixSet := redis.Redis.Scan(0,SessionStoreRedisPrefix,1000)
	for _,key := range prefixSet.Args() {
		res := redis.Redis.Del(key.(string))
		if res.Err() != nil {
			return res.Err()
		}
	}
	return nil
}
func (redis *sessionStorage) Close() error {
	return redis.Redis.Close()
}

func SessionStorage() {
	MySessionStore = session.New(session.Config{
		Expiration: 24 * time.Hour,
		Storage: &sessionStorage{Redis: storage.Redis},
		KeyLookup: "cookie:k8s_cluster_api_cookie",
		CookieDomain: envConfig.Config("BACKEND_COOKIE_DOMAIN"),
		CookiePath: "/",
		//CookieHTTPOnly: true,
		//CookieSecure: true,
		CookieSameSite: "Lex",
		KeyGenerator: utils.UUID,
	})
}