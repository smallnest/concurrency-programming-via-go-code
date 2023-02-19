package main

import "sync"

var s string
var wg sync.WaitGroup

func foo() {
	s = "hello, world"
	wg.Done() // ①
}

func main() {
	wg.Add(1)
	go foo()
	wg.Wait() // ②
	print(s)  // ③
}
