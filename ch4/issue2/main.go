package main

import (
	"fmt"
	"sync"
)

type TestStruct struct {
	wg sync.WaitGroup
}

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	t := &TestStruct{
		wg: w,
	}

	t.wg.Done()
	fmt.Println("finished")
}
