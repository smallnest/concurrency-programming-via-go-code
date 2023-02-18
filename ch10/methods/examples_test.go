package methods

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var x uint64 = 0
	newXValue := atomic.AddUint64(&x, 100) // newXValue == 100
	assert.Equal(t, uint64(100), newXValue)

	newXValue = atomic.AddUint64(&x, ^uint64(0)) // newXValue == 99
	assert.Equal(t, uint64(99), newXValue)

	atomic.AddUint64(&x, ^uint64(10-1)) // x == 89
	assert.Equal(t, uint64(89), x)

}

func TestCAS(t *testing.T) {
	var x uint64 = 0
	ok := atomic.CompareAndSwapUint64(&x, 0, 100) // ok == true
	assert.Equal(t, true, ok)

	ok = atomic.CompareAndSwapUint64(&x, 0, 100) // ok == false, x的原有的值不是0
	assert.Equal(t, false, ok)
}

func TestSwap(t *testing.T) {
	var x uint64 = 0
	old := atomic.SwapUint64(&x, 100) // old == 0
	assert.Equal(t, uint64(0), old)

	old = atomic.SwapUint64(&x, 100) // old == 100
	assert.Equal(t, uint64(100), old)
}

func TestLoad(t *testing.T) {
	var x uint64 = 0
	v := atomic.LoadUint64(&x) // v == 0
	assert.Equal(t, uint64(0), v)

	x = 100
	v = atomic.LoadUint64(&x) // v == 100
	assert.Equal(t, uint64(0), v)

}

func TestStore(t *testing.T) {
	var x uint64 = 0
	atomic.StoreUint64(&x, 100) // x == 100
	assert.Equal(t, uint64(100), x)
}
