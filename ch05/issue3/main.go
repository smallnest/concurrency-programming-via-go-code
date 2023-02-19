package main

import (
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	go func() {
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
	}()

	time.Sleep(time.Second)
	c2 := *c
	c2.Signal()
}
