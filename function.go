package main

import "fmt"

func ex2(input int) (int, bool) {
  return input/2, input % 2 == 0
}

func varadicFn(args ...int) int {
  sum := 0
  for _, v := range args {
    sum += v
  }

  return sum
}

func makeOddGenerator() func() int {
  i := 1

  return func() (r int) {
    r = i
    i += 2
    return
  }
}

func first() {
  fmt.Println("first")
}

func second() {
  fmt.Println("second")
}

func deferFn() {
  defer second()
  first()
}

func panicAndRecover() {
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}

func fib(x int) int {
  if x <= 1 {
    return x
  }

  return fib(x - 1) + fib(x - 2)
}

func main() {
  fmt.Println("Ex 1:")
  fmt.Println("func sum(arrs []float64) float64 {}")

  fmt.Println("Ex 2:")
  fmt.Println(ex2(1)) // (0, false)
  fmt.Println(ex2(2)) // (1, true)
  
  fmt.Println("Ex 3: 1, 2, 3")
  fmt.Println(varadicFn(1, 2, 3))

  fmt.Println("Ex 3: with slice")
  slice := []int{1,2,3}
  fmt.Println(varadicFn(slice...))

  fmt.Println("Ex 4:")
  maker := makeOddGenerator()
  fmt.Println(maker(), maker())

  fmt.Println("Ex 5:")

  deferFn()

  panicAndRecover()

  fmt.Println(fib(2))
  fmt.Println(fib(5)) // 0 1 1 2 3 5

}
