package main

import (
	"context"
	"fmt"
)

func foo(ctx context.Context) {
	ctx = context.WithValue(ctx, "myKey", "123")
	bar(ctx)
}

func bar(ctx context.Context) {
	ctx = context.WithValue(ctx, "myKey", true)
	fizz(ctx)
}

func fizz(ctx context.Context) {
	fmt.Println(ctx.Value("myKey"))
}

func main() {
	foo(context.Background())
}
