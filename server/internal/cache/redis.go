package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/younocode/go-vue-starter/server/config"
)

type RedisCache struct {
	Client *redis.Client
}

func newRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	// "redis://<user>:<pass>@localhost:6379/<db>"
	opt, err := redis.ParseURL(cfg.DSN())
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}

func NewRedisCache(cfg config.RedisConfig) (*RedisCache, error) {
	client, err := newRedisClient(cfg)
	if err != nil {
		return nil, err
	}
	return &RedisCache{
		Client: client,
	}, nil
}
