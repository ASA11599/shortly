package storage

import "sync"

type MemoryStore struct {
	sm *sync.Map
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		sm: &sync.Map{},
	}
}

func (ms *MemoryStore) Get(key string) string {
	v, hasKey := ms.sm.Load(key)
	vs, isString := v.(string)
	if !isString || !hasKey { return "" }
	return string(vs)
}

func (ms *MemoryStore) Set(key string, value string) {
	ms.sm.Store(key, value)
}

func (ms *MemoryStore) Close() {
	ms.sm = nil
}
