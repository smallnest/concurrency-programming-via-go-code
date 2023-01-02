package ch2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type TokenRecursiveMutex struct {
	sync.Mutex
	gentoken  int64
	token     int64
	recursion int32
}

func (m *TokenRecursiveMutex) GenToken() int64 {
	return atomic.AddInt64(&m.gentoken, 1)
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}

	m.Mutex.Lock()
	// this goroutine got the lock
	atomic.StoreInt64(&m.token, token)
	atomic.StoreInt32(&m.recursion, 1)
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}

	recursion := atomic.AddInt32(&m.recursion, -1)
	if recursion != 0 {
		return
	}

	atomic.StoreInt64(&m.token, 0)
	m.Mutex.Unlock()
}
