package main

import (
	"fmt"
	"net/http"
	"time"
)

func getFromBaidu() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Status)
}

func add(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("y can not be zero")
	}
	return x / y, nil
}

func goStatement() {
	// 并发输出
	go func() {
		fmt.Println("Hello from a goroutine!")
	}()

	// 并发访问
	go getFromBaidu()

	// go http.ListenAndServe(":8080", nil)

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println(err)
		}
	}()

	go add(1, 2) //nolint

	time.Sleep(time.Second)
}

func main() {
	// goStatement()

	// paramEvaluated()
	// paramEvaluated2()
	// paramEvaluated3()

	bench_quicksort()
}
