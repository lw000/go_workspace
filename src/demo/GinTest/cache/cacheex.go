package cache

import (
	//	"sync"
	//	"time"

	"github.com/muesli/cache2go"
)

type scheduleCacheEx struct {
	c *cache2go.CacheTable
}

func NewScheduleCacheEx() *scheduleCacheEx {
	ins := &scheduleCacheEx{
		c: cache2go.Cache("levi"),
	}
	return ins
}
