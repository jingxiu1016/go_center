package common

import "sync"

type (
	SafeMap struct {
		*sync.RWMutex
		Map map[string]interface{}
	}
)

func NewSafeMap() *SafeMap {
	return &SafeMap{Map: make(map[string]interface{})}
}

func (m *SafeMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.Map[key] = value
}
func (m *SafeMap) Get(key string) interface{} {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.Map[key]
	if !ok {
		return nil
	}
	return v
}
func (m *SafeMap) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.Map, key)
}

func (m *SafeMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.Map)
}
