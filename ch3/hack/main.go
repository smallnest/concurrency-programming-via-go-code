package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type RWMutex struct {
	sync.RWMutex
}

type m struct {
	w           sync.Mutex
	writerSem   uint32
	readerSem   uint32
	readerCount int32
	readerWait  int32
}

const rwmutexMaxReaders = 1 << 30

func (rw *RWMutex) ReaderCount() int {
	v := (*m)(unsafe.Pointer(&rw.RWMutex))
	r := atomic.LoadInt32(&v.readerCount)
	if r < 0 {
		r += rwmutexMaxReaders
	}

	return int(r)
}

func (rw *RWMutex) ReaderWait() int {
	v := (*m)(unsafe.Pointer(&rw.RWMutex))
	c := atomic.LoadInt32(&v.readerWait)

	return int(c)
}

func (rw *RWMutex) WriterCount() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&rw.RWMutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}

func main() {
	var mu RWMutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.RLock()
			time.Sleep(time.Hour)
			mu.RUnlock()
		}()
	}

	time.Sleep(time.Second)

	for i := 0; i < 50; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Hour)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)

	for i := 0; i < 50; i++ {
		go func() {
			mu.RLock()
			time.Sleep(time.Hour)
			mu.RUnlock()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("readers: ", mu.ReaderCount())
	fmt.Println("departing readers: ", mu.ReaderWait())
	fmt.Println("writer: ", mu.WriterCount())
}
