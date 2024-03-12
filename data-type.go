package main

import (
	"fmt"
  "math/bits"
  "unsafe"
  "reflect"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
)

func main() {
  fmt.Println("BASIC TYPE\n")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Architecture of system %s\n", bits.UintSize)
  var intt int
  typeof := 1
  fmt.Printf("Size of int %d bytes\n", unsafe.Sizeof(intt))
  fmt.Printf("Type of int %s\n",reflect.TypeOf(typeof) )
  fmt.Printf("Type of int %T\n", typeof)
  s := "ab£"
  fmt.Println([]byte(s))
  fmt.Println(len(s))
  fmt.Println("Loop")
  for _, v := range(s) {
    fmt.Println(v)
  }
}
