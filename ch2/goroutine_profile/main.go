package main

import (
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	go http.ListenAndServe("localhost:8080", nil)
	var count int64
	var mu sync.Mutex
	for {
		go func() {
			mu.Lock()
			defer mu.Unlock()

			count++
		}()
	}
}

// func main() {
// 	var count int64
// 	var mu sync.Mutex
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			mu.Lock()
// 			// defer mu.Unlock()

// 			count++
// 		}()
// 	}

// 	err := http.ListenAndServe("localhost:8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }
