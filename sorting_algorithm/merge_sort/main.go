package main

import "fmt"

var DUMMY = []int{10, 2, 3, 11, 1}

func main() {
	mergeSort(DUMMY, 0, len(DUMMY) - 1)

	for _, val := range DUMMY {
		fmt.Print(val, " ")
	}
}

func mergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	middle := left + int((right-left)/2)

	mergeSort(arr, left, middle)
	mergeSort(arr, middle+1, right)
	merge(arr, left, middle, right)

}

func merge(arr []int, left int, middle int, right int) {
	i, j, k := 0, 0, left
	lenL, lenR := middle-left+1, right-middle

	lArr := make([]int, lenL)
	rArr := make([]int, lenR+1)

	copy(lArr, arr[left:middle+1])
	copy(rArr, arr[middle+1:right+1])

	for i < lenL && j < lenR {
		if lArr[i] < rArr[j] {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
		}
		k++
	}

	if i < lenL {
		copy(arr[k:], lArr[i:])
	}

	if j < lenR {
		copy(arr[k:], rArr[j:])
	}
}