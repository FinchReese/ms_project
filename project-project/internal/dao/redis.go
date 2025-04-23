package dao

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"test.com/project-user/config"
)

type RedisCache struct {
	rdb *redis.Client
}

var Rc *RedisCache

func init() {
	rdb := redis.NewClient(config.AppConf.InitRedisOptions())
	Rc = &RedisCache{
		rdb: rdb,
	}
}

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}
