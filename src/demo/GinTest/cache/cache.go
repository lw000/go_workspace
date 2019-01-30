package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type scheduleCache struct {
	c *cache.Cache
}

var (
	instance *scheduleCache
	one      sync.Once
)

func SharedScheduleCache() *scheduleCache {
	one.Do(func() {
		instance = &scheduleCache{}
		instance.init()
	})
	return instance
}

func (this *scheduleCache) init() bool {
	this.c = cache.New(time.Minute*time.Duration(5), time.Minute*time.Duration(10))
	return true
}

func (this *scheduleCache) Add(key string, val interface{}) {
	if len(key) == 0 {
		return
	}

	if val == nil {
		return
	}

	this.c.Add(key, val, cache.DefaultExpiration)
}

func (this *scheduleCache) Find(key string) (interface{}, bool) {
	if len(key) == 0 {
		return nil, false
	}

	return this.c.Get(key)
}
