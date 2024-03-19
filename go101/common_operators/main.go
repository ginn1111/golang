package main

import (
	"fmt"
)

func println(args ...interface{}) {
	fmt.Println(args...)
}

type alias = struct {
	age  int
	name string
}

type alias2 = Person

type Person struct {
	age  int
	name string
}

func main() {
	var person1 alias2 = alias2{
		age:  24,
		name: "Thuan",
	}

	var person2 alias = person1

	fmt.Println(person2)
}
