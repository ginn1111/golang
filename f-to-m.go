package main

import "fmt"

func main() {
  var f float64

  fmt.Print("Enter a number of feet: ")
  fmt.Scanf("%f", &f)

  fmt.Println(f * 0.3048)
}
