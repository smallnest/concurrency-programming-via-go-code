package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// case1: expire
	log.Println("case1: expire")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	<-ctx.Done()
	log.Println("err:", ctx.Err())
	cancel()

	// case2: expire
	log.Println("case2: cancel")
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	cancel()
	<-ctx.Done()
	log.Println("err:", ctx.Err())

	log.Println("case3: parent cancel")
	pCtx, pCancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithDeadline(pCtx, time.Now().Add(5*time.Second))
	pCancel()
	<-ctx.Done()
	log.Println("err:", ctx.Err())
	cancel()

	parentDeadline()
}

func parentDeadline() {
	pCtx, pCancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer pCancel()
	ctx, cancel := context.WithDeadline(pCtx, time.Now().Add(time.Minute))
	defer cancel()

	deadline, _ := ctx.Deadline()
	fmt.Println("timeout:", time.Since(deadline)) // timeout: -5s 左右
}
