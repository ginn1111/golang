package main

import (
	"context"
	"fmt"
	"ginn/context/log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorator(handler))
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, 24, 123)
	log.Println(ctx, "handler start")

	defer log.Println(ctx, "handle ended")

	select {
	case <-time.After(time.Second * 2):
		fmt.Fprint(w, "Hello")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
