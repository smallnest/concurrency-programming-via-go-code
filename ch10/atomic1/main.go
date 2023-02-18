package main

const x int64 = 1 + 1<<33

func main() {
	var i = x
	_ = i
}
