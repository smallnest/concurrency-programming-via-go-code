package main

import (
	"sync"
)

func main() {

	var wg sync.WaitGroup
	go func() {
		for {
			wg.Add(1)
			wg.Done()
		}
	}()

	for {
		wg.Wait()
	}
}
