package db

import (
	"sync"
	"time"
)

type lruCache struct {
	timestamp time.Time
	data      []Hacker
	mu        sync.RWMutex
}

const TTL = 1

var cache lruCache = lruCache{timestamp: time.Now(), data: NewData()}

func (rdb *DB) TakeDataWithCache() ([]Hacker, error) {
	now := time.Now()

	cache.mu.RLock()
	elapsed := now.Sub(cache.timestamp)
	cache.mu.RUnlock()

	var err error = nil

	if (elapsed.Seconds() >= TTL) || (len(cache.data) == 0) {
		cache.mu.Lock()
		cache.timestamp = now
		err = rdb.loadData(&cache.data)
		cache.mu.Unlock()
	}
	return cache.data, err
}
