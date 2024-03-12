package main

import "fmt"

func main() {
  fmt.Println("Ex1: ")
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Print(i, " ")
    }
  }

  fmt.Println("\nEx:2")
  for i := 1; i <= 100; i++ {
    r := ""
    if i % 3 == 0 {
      r += "Fizz"
    }

    if i % 5 == 0 {
      r += "Buzz"
    }

    if r != "" {
      fmt.Println(i, r)
    }
  }


}
