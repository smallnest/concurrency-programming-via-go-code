package ch2

import (
	"testing"
)

func Test_reentrantFoobar(t *testing.T) {
	reentrantFoobar()
}

func Test_mutexT_foo(t *testing.T) {
	var mt mutexT
	mt.foo()
}

func Test_copyMutex(t *testing.T) {
	copyMutex()
}
