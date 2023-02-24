package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)

	for i := 0; i < 10; i++ {
		res, err := limiter.Allow(ctx, "token:123", redis_rate.PerSecond(5))
		if err != nil {
			panic(err)
		}
		log.Println("allowed", res.Allowed, "remaining", res.Remaining, "retry after", res.RetryAfter)
		if res.Allowed == 0 {
			time.Sleep(res.RetryAfter)
		}
	}

	log.Println("change the limit to 1 per second")

	for i := 0; i < 10; i++ {
		res, err := limiter.Allow(ctx, "token:123", redis_rate.PerSecond(1))
		if err != nil {
			panic(err)
		}
		log.Println("allowed", res.Allowed, "remaining", res.Remaining, "retry after", res.RetryAfter)
		if res.Allowed == 0 {
			time.Sleep(res.RetryAfter)
		}
	}
}
