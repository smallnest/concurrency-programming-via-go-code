package ch3

import (
	"sync"
	"testing"
	"time"
)

func TestReaderWriter_Priority(t *testing.T) {
	var mu sync.RWMutex

	// 制造一个departing reader
	go func() {
		mu.RLock()
		t.Log("reader1 locked")
		time.Sleep(2 * time.Second)
		mu.RUnlock()
	}()

	// writer1 请求写锁，因为有departing reader，所以writer1会等待
	go func() {
		time.Sleep(1 * time.Second)
		mu.Lock()
		t.Log("writer1 locked")
		time.Sleep(2 * time.Second)
		mu.Unlock()
	}()

	// 再有reader2请求读锁，因为有writer1在等待，所以reader2会等待
	go func() {
		time.Sleep(2 * time.Second)
		mu.RLock()
		t.Log("reader2 locked")
		time.Sleep(2 * time.Second)
		mu.RUnlock()
	}()

	// 再有writer2请求写锁，因为和writer1是互斥的，所以也会等待writer1完成后再获取到锁
	time.Sleep(1500 * time.Millisecond) // 虽然这里等待了1.5秒，比reader2先请求锁，但是还是可能reader2先获取到锁
	mu.Lock()
	t.Log("writer2 locked")
	time.Sleep(2 * time.Second)
	mu.Unlock()

}

type S struct {
	mu     sync.RWMutex
	values map[string]string
}

var s = &S{
	values: make(map[string]string),
}
