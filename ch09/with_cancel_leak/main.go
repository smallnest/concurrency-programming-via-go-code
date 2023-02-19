package main

import (
	"context"
	"log"
	"sync/atomic"
)

func main() {
	foo := func(ctx context.Context, n *int64, ch chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				log.Panicln("context is canceled")
			default:
				if atomic.AddInt64(n, 1) == 100 {
					ch <- struct{}{}
					return
				}

			}
		}
	}

	ctx, _ := context.WithCancel(context.Background())
	ch := make(chan struct{}, 1)
	n := int64(0)
	for i := 0; i < 10; i++ {
		go foo(ctx, &n, ch)
	}

	<-ch

	log.Println("n:", n)

}
