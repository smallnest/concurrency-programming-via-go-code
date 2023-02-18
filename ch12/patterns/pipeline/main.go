package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func asStream[T any](done <-chan struct{}, values ...T) <-chan T {
	s := make(chan T)
	go func() {
		defer close(s)

		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v:
			}
		}

	}()
	return s
}

func sqrt[T constraints.Integer](in <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	s := asStream(done, 1, 2, 3, 4, 5)
	out := sqrt(s)

	for v := range out {
		fmt.Println(v)
	}
}
