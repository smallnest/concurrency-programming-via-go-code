package ch3

import (
	"sync"
	"testing"
	"time"
)

func TestReentrant_Lock(t *testing.T) {
	var mu sync.RWMutex

	mu.Lock()
	{
		mu.Lock()
		t.Log("不能到达到")
		mu.Unlock()
	}
	mu.Unlock()
}

func TestReentrant_RLock(t *testing.T) {
	var mu sync.RWMutex

	mu.RLock()
	{
		mu.RLock()
		t.Log("能够达到")
		mu.RUnlock()
	}
	mu.RUnlock()
}

func TestReentrant_DeadLock(t *testing.T) {
	var mu sync.RWMutex

	// 读锁递归调用
	go func() {
		mu.RLock()
		{
			time.Sleep(10 * time.Second)
			mu.RLock()
			t.Log("不能到达到")
			mu.RUnlock()
		}
		mu.RUnlock()
	}()

	time.Sleep(1 * time.Second)

	// 在读锁递归调用前调用写锁请求
	mu.Lock()
	t.Log("不能到达到")
	mu.Unlock()
}

func TestReentrant_DeadLock2(t *testing.T) {
	var mu sync.RWMutex

	mu.RLock()
	{
		mu.Lock()
		t.Log("不能到达到")
		mu.Unlock()
	}
	mu.RUnlock()
}

func TestReentrant_DeadLock3(t *testing.T) {
	var mu sync.RWMutex

	mu.Lock()
	{
		mu.RLock()
		t.Log("不能到达到")
		mu.RUnlock()
	}
	mu.Unlock()
}

