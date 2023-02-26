package mutexmap

import (
	"sync"
)

type RWMap[K comparable, V any] struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[K]V
}

// 新建一个RWMap
func NewRWMap[K comparable, V any](n int) *RWMap[K, V] {
	return &RWMap[K, V]{
		m: make(map[K]V, n),
	}
}

func (m *RWMap[K, V]) Get(k K) (V, bool) { //从map中读取一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func (m *RWMap[K, V]) Set(k K, v V) { // 设置一个键值对
	m.Lock() // 锁保护
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap[K, V]) Delete(k K) { //删除一个键
	m.Lock() // 锁保护
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap[K, V]) Len() int { // map的长度
	m.RLock() // 锁保护
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap[K, V]) Each(f func(k K, v V) bool) { // 遍历map
	m.RLock() //遍历期间一直持有读锁
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
