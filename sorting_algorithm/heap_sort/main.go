package main

import (
	"fmt"
	"math/rand"
)

var DUMMY = []int{10, 2, 3, 11, 1, 20, 19, 15}

func main() {
 arr := DUMMY 
	heapSort(arr, len(arr))
	for _, val := range arr {
		fmt.Println(val, " ")
	}
}

func heapify(arr []int, most int, len int) {
	i := most

	l := 2*i + 1
	r := 2*i + 2

	if l < len && arr[i] < arr[l] {
		i = l
	}

	if r < len && arr[i] < arr[r] {
		i = r
	}

	if i != most {
		arr[i], arr[most] = arr[most], arr[i]
		heapify(arr, i, len)
	}
}

func heapSort(arr []int, len int) {
	for i :=  int(len/2)-1; i >= 0; i-- {
		heapify(arr, i, len)
	}

	for i := len-1; i >0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func randArr (len int) []int {

	arr := make([]int, len)

	for i := 0; i < len; i++ {
		arr[i] = rand.Intn(20)
	}

	return arr
}