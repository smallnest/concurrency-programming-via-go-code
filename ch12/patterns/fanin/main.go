package main

import (
	"fmt"
	"reflect"
	"sync"
)

func fanIn[T any](chans ...<-chan T) <-chan T {
	out := make(chan T)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan T) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()
	return out
}

func fanInReflect[T any](chans ...<-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { //remove this case
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(T)
		}
	}()
	return out

}

func fanInRec[T any](chans ...<-chan T) <-chan T {
	switch len(chans) {
	case 0:
		c := make(chan T)
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRec(chans[:m]...),
			fanInRec(chans[m:]...))
	}
}

func mergeTwo[T any](a, b <-chan T) <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func asStream(done <-chan struct{}) <-chan int {
	s := make(chan int)
	values := []int{1, 2, 3, 4, 5}
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

func main() {
	fmt.Println("fanIn by goroutine:")
	done := make(chan struct{})
	ch := fanIn(asStream(done), asStream(done), asStream(done))
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("fanIn by reflect:")
	ch = fanInReflect(asStream(done), asStream(done), asStream(done))
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("fanIn by recursion:")
	ch = fanInRec(asStream(done), asStream(done), asStream(done))
	for v := range ch {
		fmt.Println(v)
	}
}
