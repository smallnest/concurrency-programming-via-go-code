package ch2

import (
	"sync"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var mu sync.Mutex

	go func() {
		mu.Lock()
		time.Sleep(2 * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	if mu.TryLock() {
		println("try lock success")
		mu.Unlock()
	} else {
		println("try lock failed")
	}
}
