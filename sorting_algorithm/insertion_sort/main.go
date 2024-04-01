package main

import "fmt"

var DUMMY = []int{10, 3, 2, 1, 11}

func main() {
  insertionSort((DUMMY))

  for _, val := range DUMMY {
    fmt.Print(val, " ")
  }
}

func insertionSort(arr []int) {

  for i := 0; i < len(arr); i++ {

    key := arr[i]

    j := i - 1

    for j >= 0 && key < arr[j] {
      arr[j + 1] = arr[j]
      j--
    }
    arr[j+1] = key
  }
}