package main

import "fmt"

func main() {
  fmt.Println("Array")

  //
  fmt.Println("Slice")
  fmt.Println("4th elements of slice")
  slice1 := []int {1,2,3,4,5}
  fmt.Printf("Origin slice: %s\n", slice1)
  fmt.Println(slice1[:4])
  fmt.Println(len(make([]int, 3, 9)))

  fmt.Println("Find smallest in list")
  list := []int {
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }

  var smallest = list[0]

  for i := 1; i < len(list); i++ {
    if list[i] < smallest {
      smallest = list[i]
    }
  }

  fmt.Println("smallest in list", list, " is: ", smallest)

  //
  fmt.Println("Map")

}
