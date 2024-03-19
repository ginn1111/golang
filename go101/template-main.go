package main

import (
	"fmt"
)

func println(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	println("Hello go!")

}
