package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var flag atomic.Bool

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				if !flag.CompareAndSwap(false, true) { // 已经有goroutine在执行了，回避
					time.Sleep(time.Second)
					continue
				}

				// 这里做一些业务逻辑，只有一个goroutine能进来做

				flag.Store(false)
			}
		}()
	}

	wg.Wait()
}
