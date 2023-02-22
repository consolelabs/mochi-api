package cache

import (
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/response"
)

type Cache interface {
	Keys(pattern string) ([]string, error)
	Type(key string) (string, error)

	Set(key string, value interface{}, expiration time.Duration) error
	ZSet(key string, value interface{}, score float64) error

	Remove(key string) error
	ZRemove(key string, value interface{}) error
	ZRemoveByScore(key, min, max string) error
	GetString(key string) (string, error)
	GetStringSorted(key, min, max string) []string
	GetStringSortedWithScores(key, min, max string) []response.ZSetWithScoreData
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)

	HashSet(key string, hash map[string]string, expiration time.Duration) error
	HashGet(key string) (map[string]string, error)

	Publish(channel string, payload interface{}) error
	Subcribe(channel string) *redis.PubSub
}
