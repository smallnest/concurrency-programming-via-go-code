package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	testCancelCtx()
	testDeadlineCtx()
}

func testCancelCtx() {
	pCtx, cancel := context.WithCancel(context.Background())
	ctx1, _ := context.WithCancel(pCtx)

	// 撤销父context
	cancel()
	fmt.Println("cancel cancelCtx", ctx1.Err()) // context canceled
}

func testDeadlineCtx() {
	pCtx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	ctx1, _ := context.WithDeadline(pCtx, time.Now().Add(10*time.Second))

	// 撤销父context
	cancel()
	fmt.Println("cancel timerCtx", ctx1.Err()) // context deadline exceeded
}
