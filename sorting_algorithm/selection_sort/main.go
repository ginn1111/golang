package main

import (
	"fmt"
	"math"
)

var DUMMY = []int{10, 3, 2, 1, 11}
/*
	1, 3, 2, 10, 11

	1, 2, 3, 10, 11

*/

func main() {

	selectionSort(DUMMY)

	for _, val := range DUMMY {
		fmt.Print(val, " ")
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		minIdx := findMin(arr, i + 1)
		if(arr[i] > arr[minIdx]) {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}

		fmt.Println(arr)
	}

}

func findMin(arr []int, left int) int {
	minVal := math.MaxInt64
	minIdx := left;

	for i:=left; i < len(arr) ; i++ {
		if minVal > arr[i] {
			minVal = arr[i]
			minIdx = i
		}
	}

	fmt.Println(minVal, minIdx)

	return minIdx
}