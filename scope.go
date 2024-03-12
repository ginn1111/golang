package main

import  (
  "fmt"
  "reflect"

)


func test() map[string]int {
  return map[string]int{
    "a": a,
  }
}

  // func caller() {
  //   a := 2
  //   fmt.Println(test())
  // }
  //
  //
func main() {
  a := 2
  test() 
}
