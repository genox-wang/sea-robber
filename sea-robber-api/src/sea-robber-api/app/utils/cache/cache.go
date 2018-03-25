package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	// Cache go cache instance
	Cache *cache.Cache
)

const (
	// CACHE_DEFAULT_EXPIRATION  cache default expiration
	CACHE_DEFAULT_EXPIRATION = cache.DefaultExpiration
)

func init() {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
