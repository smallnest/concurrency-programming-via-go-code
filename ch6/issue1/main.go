package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var m map[int]int

func main() {
	var o Once

	initM := func() {
		time.Sleep(2 * time.Second)
		m = make(map[int]int)
	}
	go o.Do(initM) // 并发初始化

	time.Sleep(time.Second)
	o.Do(initM)
	m[1] = 1

}

type Once struct {
	done uint32
	sync.Once
}

func (o *Once) Do(f func()) {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}
