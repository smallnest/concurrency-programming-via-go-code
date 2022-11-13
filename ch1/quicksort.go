package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func quickSort_go(a []int, lo, hi int, done chan struct{}) {
	if lo >= hi {
		done <- struct{}{}
		return
	}

	p := partition(a, lo, hi)
	childDone := make(chan struct{}, 2)
	go quickSort_go(a, lo, p-1, childDone)
	go quickSort_go(a, p+1, hi, childDone)
	<-childDone
	<-childDone
	done <- struct{}{}
}

func quickSort_go2(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSort_go2(a, lo, p-1, childDone, depth)
		go quickSort_go2(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
	done <- struct{}{}
}

func bench_quicksort() {
	// 生成测试数据
	rand.Seed(time.Now().UnixNano())
	n := 10000000
	testData1, testData2, testData3 := make([]int, 0, n), make([]int, 0, n), make([]int, 0, n)
	for i := 0; i < n; i++ {
		val := rand.Intn(n * 100)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
	}

	//串行程序
	start := time.Now()
	quickSort(testData1, 0, len(testData1)-1)
	fmt.Println("串行执行: ", time.Since(start))

	// 并行程序
	done := make(chan struct{})
	start = time.Now()
	go quickSort_go(testData2, 0, len(testData2)-1, done)
	<-done
	fmt.Println("完全并发执行: ", time.Since(start))

	done = make(chan struct{})
	start = time.Now()
	go quickSort_go2(testData3, 0, len(testData3)-1, done, 3)
	<-done
	fmt.Println("优化并发执行: ", time.Since(start))
}
