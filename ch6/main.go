package main

import (
	"fmt"
	"sync"

	"github.com/carlmjohnson/syncx"
)

func main() {
	var getMoL = syncx.Once(func() int {
		fmt.Println("calculating meaning of life...")
		return 42
	})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(getMoL())
			wg.Done()
		}()
	}
	wg.Wait()
}
