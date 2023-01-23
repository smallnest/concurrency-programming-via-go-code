package ch3

import (
	"sync"
	"testing"
)

func BenchmarkCounter_Mutex(b *testing.B) {
	// 1. Declare a variable counter of type int64
	var counter int64
	// 2. Declare a variable mu of type sync.Mutex
	var mu sync.Mutex

	for i := 0; i < b.N; i++ {
		// 3. Use RunParallel to run the benchmark in parallel
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			// 4. Use pb.Next() in a for loop to iterate over the tests
			for pb.Next() {
				i++

				// 5. If i is a multiple of 10000, then increment the counter
				if i%10000 == 0 {
					mu.Lock()
					counter++
					mu.Unlock()
				} else {
					// 6. Otherwise, read the counter
					mu.Lock()
					_ = counter
					mu.Unlock()
				}

			}
		})
	}

}

func BenchmarkCounter_RWMutex(b *testing.B) {
	var counter int64
	var mu sync.RWMutex

	for i := 0; i < b.N; i++ {
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				i++

				if i%10000 == 0 {
					// Lock the mutex, increment the counter, and unlock the mutex
					mu.Lock()
					counter++
					mu.Unlock()
				} else {
					// Lock the mutex for reading, retrieve the counter value, and unlock the mutex
					mu.RLock()
					_ = counter
					mu.RUnlock()
				}

			}
		})
	}

}
