package main

import (
	"log"
	"time"

	"github.com/juju/ratelimit"
)

func main() {
	var bucket = ratelimit.NewBucket(time.Second, 3)
	for i := 0; i < 10; i++ {
		bucket.Wait(1)
		log.Printf("got #%d", i)
	}
}
