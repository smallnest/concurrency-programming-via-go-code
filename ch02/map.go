package ch2

import "sync"

type Map[k comparable, v any] struct {
	mu sync.Mutex
	m  map[k]v
}

func NewMap[k comparable, v any](size ...int) *Map[k, v] {
	if len(size) > 0 {
		return &Map[k, v]{
			m: make(map[k]v, size[0]),
		}
	}
	return &Map[k, v]{
		m: make(map[k]v),
	}
}

func (m *Map[k, v]) Get(key k) (v, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.m[key]
	return value, ok
}

func (m *Map[k, v]) Set(key k, value v) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[key] = value
}
