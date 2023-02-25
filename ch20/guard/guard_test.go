package guard

import (
	"testing"
)

func TestGuard(t *testing.T) {
	Guard(func() {
		panic("panic in guard")
	})
}

func TestGuardClose(t *testing.T) {
	ch := make(chan int)
	close(ch)
	GuardClose(ch)

	ch = nil
	GuardClose(ch)

	ch = make(chan int, 1)
	close(ch)
	GuardSend(ch, 100)
}
