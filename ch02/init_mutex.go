package ch2

import "sync"

func useMutex() {
	var mu sync.Mutex
	// mu := sync.Mutex{}
	mu.Lock()
	// do something
	mu.Unlock()
}

func useMutex2() {
	type T struct {
		mu sync.Mutex
		m  map[int]int
	}

	var t = &T{
		// mu: sync.Mutex{},
		m: make(map[int]int),
	}

	t.mu.Lock()
	// do something
	t.mu.Unlock()
}
