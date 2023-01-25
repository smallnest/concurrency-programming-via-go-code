package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelta(t *testing.T) {
	var counter uint64 = 100
	var delta int = -1
	fmt.Printf("%b\n%b\n", counter, uint64(delta))
	counter += uint64(delta)

	assert.Equal(t, uint64(99), counter)

}
