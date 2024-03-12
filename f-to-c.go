package main

import "fmt"

func main() {

  var f float64
  fmt.Print("Enter a fahrenheit number: ")

  fmt.Scanf("%f", &f)

  fmt.Println((f - 32) * 5/9)
}
