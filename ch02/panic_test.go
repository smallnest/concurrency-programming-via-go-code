package ch2

import (
	"sync"
	"testing"
)

func TestOnlyUnlock(t *testing.T) {
	var mu sync.Mutex
	mu.Unlock()
}
