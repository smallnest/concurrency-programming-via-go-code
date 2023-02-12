package main

import (
	"context"
	"fmt"

	xcontext "golang.org/x/net/context"
)

func foobar(ctx context.Context) {
	fmt.Println("define as context.Context but use xcontext.Context")
}

func main() {
	var ctx xcontext.Context = xcontext.Background()
	foobar(ctx)
}
