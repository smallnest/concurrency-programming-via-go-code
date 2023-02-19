package main

import (
	"fmt"
	"time"
)

func main() {
	var list []int

	go func(l int) {
		time.Sleep(time.Second)
		fmt.Printf("passed len: %d, current list len:%d\n", l, len(list))
	}(len(list))

	list = append(list, 1)

	time.Sleep(2 * time.Second)

}
