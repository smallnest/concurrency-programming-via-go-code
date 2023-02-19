package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		srv := ctx.Value(http.ServerContextKey).(*http.Server)
		fmt.Printf("server: %v\n", srv.Addr)

		local := ctx.Value(http.LocalAddrContextKey).(net.Addr)
		fmt.Printf("local: %s\n", local)
	})

	go http.ListenAndServe(":8080", nil)

	resp, _ := http.Get("http://localhost:8080/")
	resp.Body.Close()
}
