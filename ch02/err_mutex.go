package ch2

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu      sync.Mutex
	counter int64
}

func (c *Counter) Inc() {
	c.mu.Lock()
	// defer c.mu.Unlock()
	c.counter++
}

func reentrantFoobar() {
	var count int64
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	mu.Lock()
	count++
	mu.Unlock()
}

type mutexT struct {
	mu sync.Mutex
}

func (t *mutexT) foo() {
	t.mu.Lock()
	defer t.mu.Unlock()

	fmt.Println("in bar")

	t.bar()

}
func (t *mutexT) bar() {
	t.mu.Lock()
	defer t.mu.Unlock()

	fmt.Println("in bar")
}

func copyMutex() {
	var mu sync.Mutex
	var mu2 sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	mu2 = mu

	mu2.Lock()
	// do something
	mu2.Unlock()
}
