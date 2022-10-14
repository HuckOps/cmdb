package db

import (
	"cmdb/pkg/config"
	"fmt"
	"github.com/go-redis/redis"
)

var Redis RedisClient
type RedisClient struct {
	TokenCache *redis.Client
}
func InitRedis()  {
	 Redis.TokenCache = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%d", config.Config.Redis.Host,
		config.Config.Redis.Port), DB: 0})
	if _, err := Redis.TokenCache.Ping().Result(); err != nil {
		panic(err)
	}
}