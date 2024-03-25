package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type Key int

var key = Key(24)

func Println(ctx context.Context, msg string) {
	ctxValue, ok := ctx.Value(key).(int64)

	if !ok {
		log.Print("could not find request Id in context")
	}

	log.Printf("[%d] %v", ctxValue, msg)
}

func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := rand.Int63()
		ctx := context.WithValue(r.Context(), key, id)

		f(w, r.WithContext(ctx))
	}
}
