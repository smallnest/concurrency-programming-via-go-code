package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.WithValue(context.Background(), "key1", "0001")
	ctx2 := context.WithValue(ctx1, "key2", "0002")
	ctx3 := context.WithValue(ctx2, "key3", "0003")
	ctx4 := context.WithValue(ctx3, "key4", "0004")

	fmt.Println(ctx4.Value("key1"))
}
