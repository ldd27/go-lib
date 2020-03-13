package redistool

import (
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
)

type RedisConf struct {
	Address  string
	Password string
	DB       int
}

var defaultRedisConf = RedisConf{
	DB: 0,
}

func NewRedis(opts ...func(*RedisConf)) *redis.Client {
	conf := defaultRedisConf
	for _, o := range opts {
		o(&conf)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.DB,
	})

	return client
}

func NewRedisRateLimiter(client *redis.Client) *redis_rate.Limiter {
	return redis_rate.NewLimiter(client)
}
