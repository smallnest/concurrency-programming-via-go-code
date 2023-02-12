package main

import (
	"context"
	"fmt"
	"io"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(io.EOF)

	fmt.Println(ctx.Err())          // 返回 context.Canceled
	fmt.Println(context.Cause(ctx)) // 返回 io.EOF

	ctx, cancel = context.WithCancelCause(context.Background())
	cancel(nil)

	fmt.Println(ctx.Err())          // 返回 context.Canceled
	fmt.Println(context.Cause(ctx)) // 返回 context.Canceled
}
