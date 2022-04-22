package cache

import (
	"time"
)

type Cache interface {
	Keys(pattern string) ([]string, error)
	Type(key string) (string, error)

	Set(key string, value interface{}, expiration time.Duration) error
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)

	HashSet(key string, hash map[string]string, expiration time.Duration) error
	HashGet(key string) (map[string]string, error)
}
