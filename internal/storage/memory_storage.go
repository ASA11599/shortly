package storage

type MemoryStorage struct {
	m map[string]string
}

var memoryStorage *MemoryStorage

func GetMemoryStorage() (*MemoryStorage, error) {
	if memoryStorage == nil {
		memoryStorage = &MemoryStorage{
			m: make(map[string]string),
		}
		return memoryStorage, nil
	}
	return memoryStorage, nil
}

func (ms *MemoryStorage) Get(key string) (string, error) {
	return ms.m[key], nil
}

func (ms *MemoryStorage) Set(key string, value string) error {
	ms.m[key] = value
	return nil
}

func (rs *MemoryStorage) Close() error {
	memoryStorage = nil
	return nil
}
