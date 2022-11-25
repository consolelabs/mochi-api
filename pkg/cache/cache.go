package cache

import (
	"time"
)

type Cache interface {
	Keys(pattern string) ([]string, error)
	Type(key string) (string, error)

	Set(key string, value interface{}, expiration time.Duration) error
	ZSet(key string, value interface{}, score float64) error

	Remove(key string) error
	ZRemove(key string, value interface{}) error
	GetString(key string) (string, error)
	GetStringSorted(key, min, max string) []string
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)

	HashSet(key string, hash map[string]string, expiration time.Duration) error
	HashGet(key string) (map[string]string, error)
}
