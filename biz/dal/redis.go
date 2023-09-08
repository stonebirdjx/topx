package dal

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() {
	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		panic(err)
	}

	_ = redis.NewClient(opt)
}
