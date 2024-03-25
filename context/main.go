package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	sleepAndTalk(ctx, time.Second*5, "Hello")
}

func sleepAndTalk(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
