package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	reserve()
}

func quickstart() {
	var limit = rate.Every(200 * time.Millisecond)
	var limiter = rate.NewLimiter(limit, 3)
	for i := 0; i < 15; i++ {
		log.Printf("got #%d, err:%v", i, limiter.Wait(context.Background()))
	}
}

func setLimit() {

	var limiter = rate.NewLimiter(1, 3)
	for i := 0; i < 3; i++ {
		log.Printf("got #%d, err:%v", i, limiter.Wait(context.Background()))
	}

	log.Println("set new limit at 10s")
	limiter.SetLimitAt(time.Now().Add(10*time.Second), rate.Every(3*time.Second))

	for i := 4; i < 9; i++ {
		log.Printf("got #%d, err:%v", i, limiter.Wait(context.Background()))
	}
}

func reserve() {
	var limiter = rate.NewLimiter(1, 10)
	limiter.WaitN(context.Background(), 10) // 把初始的令牌消耗掉

	r := limiter.ReserveN(time.Now().Add(5), 4)
	log.Printf("ok: %v, delay: %v", r.OK(), r.Delay())
	r.Cancel()
	r = limiter.ReserveN(time.Now().Add(3), 6)
	log.Printf("ok: %v, delay: %v", r.OK(), r.Delay())
	r = limiter.ReserveN(time.Now().Add(3), 100)
	log.Printf("ok: %v, delay: %t", r.OK(), r.Delay() == rate.InfDuration)
}
