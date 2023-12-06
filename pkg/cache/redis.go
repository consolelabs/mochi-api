package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/response"
)

type redisCache struct {
	rdb *redis.Client
	ctx context.Context
}

type RedisOpts struct {
	URL          string
	SentinelURLs []string
	MasterName   string
}

func NewRedisCache(opts RedisOpts) (Cache, error) {
	var rdb *redis.Client
	if opts.MasterName == "" {
		redisURL := opts.URL
		if redisURL == "" {
			return nil, fmt.Errorf("redis url is not set")
		}
		redisOpt, err := redis.ParseURL(redisURL)
		if err != nil {
			return nil, err
		}
		rdb = redis.NewClient(redisOpt)
	} else {
		if len(opts.SentinelURLs) == 0 {
			return nil, fmt.Errorf("redis sentinel url is not set")
		}
		rdb = redis.NewFailoverClient(&redis.FailoverOptions{
			SentinelAddrs: opts.SentinelURLs,
			MasterName:    opts.MasterName,
		})
	}
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

func (c *redisCache) ZSet(key string, value interface{}, score float64) error {
	return c.rdb.ZAdd(context.Background(), key, &redis.Z{
		Score:  score,
		Member: value,
	}).Err()
}

func (c *redisCache) Remove(key string) error {
	return c.rdb.Del(context.Background(), key).Err()
}

func (c *redisCache) ZRemove(key string, value interface{}) error {
	return c.rdb.ZRem(context.Background(), key, value).Err()
}

func (c *redisCache) ZRemoveByScore(key, min, max string) error {
	return c.rdb.ZRemRangeByScore(context.Background(), key, min, max).Err()
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

func (c *redisCache) GetStringSorted(key, min, max string) []string {
	return c.rdb.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Val()
}

func (c *redisCache) GetStringSortedWithScores(key, min, max string) []response.ZSetWithScoreData {
	results := c.rdb.ZRangeByScoreWithScores(context.Background(), key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Val()
	parsedResults := make([]response.ZSetWithScoreData, 0, len(results))
	for _, v := range results {
		value := response.ZSetWithScoreData{}
		value.Member = fmt.Sprintf("%v", v.Member)
		value.Score = v.Score
		parsedResults = append(parsedResults, value)
	}
	return parsedResults
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
	err := c.rdb.HSet(context.Background(), key, hash).Err()
	if err != nil {
		return err
	}
	c.rdb.ExpireAt(context.Background(), key, time.Now().Add(expiration))
	return err
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

func (c *redisCache) Publish(channel string, payload interface{}) error {
	err := c.rdb.Publish(context.Background(), channel, payload).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *redisCache) Subcribe(channel string) *redis.PubSub {
	subcriber := c.rdb.Subscribe(context.Background(), channel)
	return subcriber
}
