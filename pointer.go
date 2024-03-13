package main

import "fmt"

func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func square(n *float64) {
  *n = *n * *n
}

func main() {
  var a int
  fmt.Println("get the memory address using & sign", &a)
  fmt.Println("get the value using * sign", *(&a))
  fmt.Println("create new pointer using new operator", new(int))

  x := 1.5
  square(&x)
  fmt.Println(x)

  xx := 1
  yy := 2

  fmt.Printf("xx: %d, yy: %d", xx, yy)
  swap(&xx, &yy)
  fmt.Printf("\nAfter swap, xx: %d, yy: %d", xx, yy)
}
