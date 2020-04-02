/*
@Time : 2020/4/2 9:45 AM
*/
package cache

import (
	"sync"

	"github.com/xormplus/core"
)

type MemoryCache struct {
	store map[interface{}]interface{}
	mu    sync.Mutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{store: make(map[interface{}]interface{})}
}

func (m MemoryCache) Put(key string, value interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[key] = value
	return nil
}

func (m MemoryCache) Get(key string) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.store[key]; ok {
		return v, nil
	}
	return nil, core.ErrCacheMiss
}

func (m MemoryCache) Del(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
	return nil
}

var _ core.CacheStore = (*MemoryCache)(nil)
