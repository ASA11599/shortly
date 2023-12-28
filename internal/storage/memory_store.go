package storage

import "sync"

type MemoryStore struct {
	m map[string]string
	lock sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		m: make(map[string]string),
	}
}

func (ms *MemoryStore) Get(key string) string {
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	return ms.m[key]
}

func (ms *MemoryStore) Set(key string, value string) {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	ms.m[key] = value
}

func (ms *MemoryStore) Close() {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	clear(ms.m)
	ms.m = nil
}
