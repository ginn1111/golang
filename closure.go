package main

import "fmt"

func closureFactory() func() (int) {
  i := 0
  return func() (r int) {
    r = i
    i += 1

    return
  }
}

func main() {
  a := closureFactory()

  fmt.Println(a())
  fmt.Println(a())

}
