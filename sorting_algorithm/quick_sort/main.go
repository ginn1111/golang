package main

import "fmt"

var DUMMY = []int{-1, 0, 1, 2, -1, 4}

func main() {
	qSort(DUMMY, 0, len(DUMMY) - 1)

	for _, val := range DUMMY {
		fmt.Print(val, " ")
	}
}

func partitionLeft(arr []int, left, right int) int {
	i := right + 1
	
	for j:=right; j > left; j-- {
		if arr[left] < arr[j] {
			i--
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i-1], arr[left] = arr[left], arr[i-1]
	return i-1

}

func partitionRight(arr []int, left, right int) int {
	i := left - 1
	p := arr[right]

	for j := left; j < right; j++ {
		if arr[j] < p {
			i++;
			arr[i], arr[j] = arr[j], arr[i]
		}

	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	fmt.Println(i)

	return i+1
}

func qSort(arr []int, left, right int) {
	if left < right {

		p := partitionRight(arr, left, right)

		qSort(arr, left, p - 1)
		qSort(arr, p + 1, right)
	}
}

func qSortMid(arr []int, left int, right int) {
	if left >= right {
		return
	}
	p := left
	pVal := arr[p]
	i, j := left + 1, right

	for i < j {
		for arr[i] < pVal {
			i++
		}

		for arr[j] > pVal {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}

	}

		qSortMid(arr, left, p)
		qSortMid(arr, p+1, right)
}
