package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	ch := make(chan int)
	handleTimeout(ch, time.Second)

	handleTimeAfter(ch, time.Second)

	conn, err := net.DialTimeout("tcp", "127.0.0.1", time.Second)

	httpClient := &http.Client{
		Timeout: time.Second,
	}

	_, _ = conn, err
	_ = httpClient
}

func handleTimeout(readCh chan int, maxTime time.Duration) {
	timer := time.NewTimer(maxTime)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			println("timeout")
			return
		case count := <-readCh:
			if count == 100 {
				println("done")
			}
		}
	}
}

func handleTimeAfter(readCh chan int, maxTime time.Duration) {
	for {
		select {
		case <-time.After(maxTime):
			println("timeout")
			return
		case count := <-readCh:
			if count == 100 {
				println("done")
			}
		}
	}
}
