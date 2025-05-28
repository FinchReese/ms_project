package dao

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Rdb *redis.Client
}

var Rc *RedisCache

func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := rc.Rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.Rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rc *RedisCache) SetAdd(ctx context.Context, key string, value string) error {
	err := rc.Rdb.SAdd(ctx, key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisCache) GetMembers(ctx context.Context, key string) ([]string, error) {
	members, err := rc.Rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (rc *RedisCache) ClearSet(ctx context.Context, key string) error {
	err := rc.Rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisCache) Delete(ctx context.Context, key string) error {
	err := rc.Rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
