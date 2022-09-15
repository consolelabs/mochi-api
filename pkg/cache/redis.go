package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisCache(opts *redis.Options) (Cache, error) {
	rdb := redis.NewClient(opts)
	ctx := context.Background()

	// ping to test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &redisCache{
		rdb: rdb,
	}, nil
}

func (c *redisCache) Keys(pattern string) ([]string, error) {
	res, err := c.rdb.Keys(context.Background(), pattern).Result()
	switch err {
	case nil:
		return res, nil
	case redis.Nil:
		return nil, nil
	default:
		return nil, err
	}
}

func (c *redisCache) Type(key string) (string, error) {
	res, err := c.rdb.Type(context.Background(), key).Result()
	switch err {
	case nil:
		return res, nil
	case redis.Nil:
		return "", nil
	default:
		return "", err
	}
}

func (c *redisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.rdb.Set(context.Background(), key, value, expiration).Err()
}

func (c *redisCache) Remove(key string) error {
	return c.rdb.Del(context.Background(), key).Err()
}

func (c *redisCache) GetString(key string) (string, error) {
	res, err := c.rdb.Get(context.Background(), key).Result()
	switch err {
	case nil:
		return res, nil
	case redis.Nil:
		return "", nil
	default:
		return "", err
	}
}

func (c *redisCache) GetInt(key string) (int, error) {
	res, err := c.rdb.Get(context.Background(), key).Int()
	switch err {
	case nil:
		return res, nil
	case redis.Nil:
		return 0, nil
	default:
		return 0, err
	}
}

func (c *redisCache) GetBool(key string) (bool, error) {
	res, err := c.rdb.Get(context.Background(), key).Bool()
	switch err {
	case nil:
		return res, nil
	case redis.Nil:
		return false, nil
	default:
		return false, err
	}
}

func (c *redisCache) HashSet(key string, hash map[string]string, expiration time.Duration) error {
	return c.rdb.HSet(context.Background(), key, hash).Err()
}

func (c *redisCache) HashGet(key string) (map[string]string, error) {
	res, err := c.rdb.HGetAll(context.Background(), key).Result()
	switch err {
	case nil:
		return res, nil
	case redis.Nil:
		return nil, nil
	default:
		return nil, err
	}
}
