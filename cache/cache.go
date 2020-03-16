package cache

import (
	"runtime/debug"

	"github.com/coocood/freecache"
)

var (
	client *freecache.Cache
)

// Client return client
func Client() *freecache.Cache {
	return client
}

// Init 初始化Cache
func Init(cacheSize int) {
	if cacheSize <= 0 {
		cacheSize = 10 * 1024 * 1024
	}
	if cacheSize >= 1024*1024*1024 {
		cacheSize = 1024 * 1024 * 1024
	}
	if cacheSize > 50*1024*1024 {
		debug.SetGCPercent(20)
	}
	client = freecache.NewCache(cacheSize)
}
