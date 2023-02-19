package ch2

import (
	"testing"
)

func TestBuiltinMap(t *testing.T) {
	var m = make(map[int]int)
	go func() {
		i := 0
		for {
			m[i]++
			i++
		}
	}()

	i := 0
	for {
		_ = m[i]
		i++
	}

}

// func TestNewMap(t *testing.T) {
// 	var m = NewMap[int, int]()
// 	go func() {
// 		i := 0
// 		for {
// 			v, _ := m.Get(i)
// 			m.Set(i, v+1)
// 			i++
// 		}
// 	}()

// 	i := 0
// 	for {
// 		_, _ = m.Get(i)
// 		i++
// 	}
// }
