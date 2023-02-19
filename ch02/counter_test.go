package ch2

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000000; i++ {
				counter++
			}

			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 64000000 {
		t.Errorf("counter should be 64000000, but got %d", counter)
	}
}

func TestCounterWithMutex(t *testing.T) {
	var counter int64

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000000; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	if counter != 64000000 {
		t.Errorf("counter should be 64000000, but got %d", counter)
	}
}
