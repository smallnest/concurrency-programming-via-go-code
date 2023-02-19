package ch2

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/kortschak/goroutine"
)

// RecursiveLock aka. ReentrantLock
type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int64
}

func (m *RecursiveMutex) Lock() {
	gid := goroutine.ID()
	if atomic.LoadInt64(&m.owner) == gid {
		atomic.AddInt64(&m.recursion, 1)
		return
	}
	m.Mutex.Lock()

	// this goroutine got the lock
	atomic.StoreInt64(&m.owner, gid)
	atomic.StoreInt64(&m.recursion, 1)
}

func (m *RecursiveMutex) Unlock() {
	gid := goroutine.ID()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	recursion := atomic.AddInt64(&m.recursion, -1)
	if recursion > 0 {
		return
	}

	// this goroutine releases the lock
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
