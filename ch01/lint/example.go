package main

import (
	"fmt"
	"net/http"
	"time"
)

func add(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("y can not be zero")
	}
	return x / y, nil
}

func main() {
	go http.ListenAndServe(":8080", nil)

	go add(1, 2) // nolint

	time.Sleep(time.Second)
}
