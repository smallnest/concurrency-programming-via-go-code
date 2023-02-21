package main

import (
	"context"
	"fmt"

	"github.com/smallnest/concurrency-programming-via-go-code/ch09/key3/p1"
	"github.com/smallnest/concurrency-programming-via-go-code/ch09/key3/p2"
)

func main() {
	ctx := context.WithValue(context.Background(), p1.Mykey1, "123")
	ctx = context.WithValue(ctx, p2.Mykey1, true)
	fmt.Println(ctx.Value(p1.Mykey1))
	fmt.Println(ctx.Value(p2.Mykey1))
}
