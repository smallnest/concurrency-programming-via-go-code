package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/marusama/cyclicbarrier"
)

func main() {
	cnt := 0
	b := cyclicbarrier.NewWithAction(10, func() error {
		cnt++
		return nil
	})

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				log.Printf("goroutine %d 来到第%d轮屏障", i, j)
				err := b.Await(context.TODO())
				log.Printf("goroutine %d 冲破第%d轮屏障", i, j)
				if err != nil {
					panic(err)
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println(cnt)
}
