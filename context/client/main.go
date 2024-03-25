package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// create a context
	ctx := context.Background()

	// create a timeout context with deadline is 1s
	ctx, cancel := context.WithTimeout(ctx, time.Second*7)

	var _ = cancel

	// initial http get request with context, method, url and body
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)

	// call immediately
	// defer cancel()

	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
