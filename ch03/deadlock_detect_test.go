package ch3

import (
	"testing"
	"time"

	sync "github.com/sasha-s/go-deadlock"
)

func TestReentrant_DeadLock1_Detector(t *testing.T) {
	var mu sync.RWMutex

	mu.Lock()
	{
		mu.Lock()
		t.Log("不能到达到")
		mu.Unlock()
	}
	mu.Unlock()
}

func TestReentrant_DeadLock2_Detector(t *testing.T) {
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

func TestReentrant_DeadLock3_Detector(t *testing.T) {
	var mu sync.RWMutex

	mu.RLock()
	{
		mu.Lock()
		t.Log("不能到达到")
		mu.Unlock()
	}
	mu.RUnlock()
}

func TestReentrant_DeadLock4_Detector(t *testing.T) {
	var mu sync.RWMutex

	mu.Lock()
	{
		mu.RLock()
		t.Log("不能到达到")
		mu.RUnlock()
	}
	mu.Unlock()
}
