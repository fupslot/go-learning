package cache

import (
	"path"
	"time"

	inmem "github.com/patrickmn/go-cache"
)

var (
	defaultExpire = 5 * time.Minute
	defaultPurge  = 30 * time.Second

	CachePrefix = "cache"
)

var Cache = inmem.New(defaultExpire, defaultPurge)

func BuildCacheKeys(keys ...string) string {
	keys = append([]string{CachePrefix}, keys...)
	return path.Join(keys...)
}
