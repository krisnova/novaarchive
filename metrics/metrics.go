package metrics

import (
	"fmt"
	"sync"
)

type MetricResource struct {
	backend Backend
	mutex   sync.Mutex
}

// Record is a key/value pair
// that is intentionally a type string.
type Record struct {
	Key   string
	Value string
}

func (m *MetricResource) Get(key string) (*Record, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	value, err := m.backend.Get(key)
	if err != nil {
		return nil, fmt.Errorf("unable to find record from [%s]: %v", m.backend.Type(), err)
	}
	return &Record{
		Key:   key,
		Value: value,
	}, nil
}

func (m *MetricResource) Set(key string, value string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.backend.Set(key, value)
}
