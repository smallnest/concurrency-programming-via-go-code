package ch2

import (
	"sync"
	"testing"
)

func TestBuiltinSlice(t *testing.T) {
	var s []int

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 1_000_000; i++ {
				s = append(s, 1)
			}
		}()
	}

	wg.Wait()

	if len(s) != 10*1_000_000 {
		t.Fatalf("len(s) = %d, want %d", len(s), 10*1_000_000)
	}
}

func TestSafeSlice(t *testing.T) {
	var mu sync.Mutex
	var s []int

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 1_000_000; i++ {
				mu.Lock()
				s = append(s, 1)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	if len(s) != 10*1_000_000 {
		t.Fatalf("len(s) = %d, want %d", len(s), 10*1_000_000)
	}
}
