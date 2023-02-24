package main

import (
	"log"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(1, ratelimit.WithSlack(3)) // one per second, slack=3

	for i := 0; i < 10; i++ {
		rl.Take()
		log.Printf("got #%d", i)
		if i == 3 {
			time.Sleep(5 * time.Second)
		}
	}

}
