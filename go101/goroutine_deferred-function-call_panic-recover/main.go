package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	wg.Add(2)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)

	log.Println(runtime.NumCPU()) // logical of CPUs
	log.Println(runtime.GOMAXPROCS)
	wg.Wait()
}
