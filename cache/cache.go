package cacheHelper

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var gcache *cache.Cache
var gonce sync.Once

//获取缓存对象
func GetCache() *cache.Cache {
	gonce.Do(func() {
		gcache = cache.New(60*time.Second, 10*time.Minute)
	})
	return gcache
}