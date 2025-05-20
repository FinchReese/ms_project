package config

import (
	"github.com/go-redis/redis/v8"
	"test.com/project-project/internal/dao"
)

func InitRedisClient(options *redis.Options) {
	rdb := redis.NewClient(options)
	dao.Rc = &dao.RedisCache{
		Rdb: rdb,
	}
}
