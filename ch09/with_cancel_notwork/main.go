package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {

		select {
		case <-ctx.Done():
			return
		default:
			for {
				// 一段长时间运行，无法中途终止的代码
			}
		}

	}()

	cancel()

}
