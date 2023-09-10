package dal

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	RedisRDB     *redis.Client
	RedisLimiter *redis_rate.Limiter
)

type RedisOption struct {
	URI string // "redis://<user>:<pass>@localhost:6379/<db>"
}

func RedisInit(opt RedisOption) error {
	o, err := redis.ParseURL(opt.URI)
	if err != nil {
		return err
	}

	RedisRDB = redis.NewClient(o)

	RedisLimiter = redis_rate.NewLimiter(RedisRDB)
	return nil
}
