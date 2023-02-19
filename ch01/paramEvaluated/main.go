package main

import (
	"fmt"
	"time"
)

func paramEvaluated() {
	var list []int

	go func(l int) {
		time.Sleep(time.Second)
		fmt.Printf("passed len: %d, current list len:%d\n", l, len(list))
	}(len(list))

	list = append(list, 1)

	time.Sleep(2 * time.Second)
}

func paramEvaluated2() {
	list := []int{1}

	foo := func(l int) {
		time.Sleep(time.Second)
		fmt.Printf("passed len: %d, current list len:%d\n", l, len(list))
	}

	go foo(len(list))

	foo = func(l int) {
		fmt.Printf("passed len: %d, current list len:%d\n", l*100, len(list)*100)
	}

	time.Sleep(2 * time.Second)

	foo(len(list))
}

func paramEvaluated3() {
	type Student struct {
		Name string
	}

	s := &Student{
		Name: "博文",
	}

	go func(s *Student) {
		time.Sleep(time.Second)
		fmt.Printf("student name: %s\n", s.Name)
	}(s)

	s.Name = "约礼"
	time.Sleep(2 * time.Second)
}

func main() {
	// paramEvaluated()
	// paramEvaluated2()
	paramEvaluated3()
}
