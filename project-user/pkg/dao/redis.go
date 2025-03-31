package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"test.com/project-user/config"
	"time"
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

func (rc *RedisCache) Put(key, value string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisCache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := rc.rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}
