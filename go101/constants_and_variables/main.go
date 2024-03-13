package main

import (
	"fmt"
	"reflect"
)

var x, y = a + 1, 5           // 8 5
var a, b, c = b + 1, c + 1, y // 7 6 5

const n = 1 << 64

type Test struct {
	testField int
}

func main() {
	test := Test{1}
	test.testField = 2
	fmt.Println(test)

	fmt.Println("Hello go")
	// _ := 100
	var lang, website string

	var _, _ = lang, website
	untyped_float_constant := 1.23

	float := 2e307

	fmt.Println(string(100))

	fmt.Println(uint(1.0), int16(6+0i), complex128(789))
	fmt.Println(reflect.TypeOf(untyped_float_constant))
	fmt.Println(int(untyped_float_constant), 1.23e2)
	fmt.Println(int32(float))

	fmt.Println(x > 1)
	fmt.Println(x, y, a, b, c)
}
