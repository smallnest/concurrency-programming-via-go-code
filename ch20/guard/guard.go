package guard

func Guard(fn func()) {
	defer func() {
		recover()
	}()

	fn()
}

func GuardClose[T any](ch chan T) {
	defer func() {
		recover()
	}()

	close(ch)
}

func GuardSend[T any](ch chan T, v T) {
	defer func() {
		recover()
	}()

	ch <- v
}
