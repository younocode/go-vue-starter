package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/younocode/go-vue-starter/server/config"
)

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	// "redis://<user>:<pass>@localhost:6379/<db>"
	opt, err := redis.ParseURL(cfg.DSN())
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
