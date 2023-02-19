package main

func main() {
	ch := make(chan int)

	select {
	case ch <- 1:
	default:
	}

	select {
	case <-ch:
	default:
	}
}
