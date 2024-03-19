package main

import (
	"fmt"
)

func println(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	// Retrieve and modify container
	println([...]bool{false, true, 3:true, false})
	println([...]bool{3: false, 1: true, true})
	var nilMap map[string]int

	// Container assignment
	println(nilMap["nil"]) // print zero value of element type -> print 0
	// nilMap["nil"] = 1  // will cause panic

	slice := make([]int, 3, 5)
	newSlice := append(slice, 1, 2)

	// Append and delete container
	slice[0] = 1
	fmt.Println(slice, newSlice)

	// Length of array and array pointer is typed (int) constant

	arr := [...]int{1,2,3}
	arrPointer := &[...]int{1,2,3}
	const lenArrPointer = len(arrPointer)
	const lenArr = len(arr)

	// Slice manipulations

	sliceManipulations()
}

func sliceManipulations() {
	fmt.Println("Slice Manipulations")
	fmt.Println(">Clone")

	s := []int{1,2,3}

	sClone1 := append(s[:0:0], s...)

	fmt.Println(">>Clone1")
	fmt.Println(sClone1, sClone1 == nil)
	printAddressOfPointer(s, sClone1)

	fmt.Println(">>Clone2")
	s = []int{}
	sClone2 := append([]int(nil), s...)
	fmt.Println(sClone2 == nil)
	
}

func printAddressOfPointer(args...interface{}) {
	var repeatStr string
	for  range args {
		repeatStr += "%p "
	}
	repeatStr += "\n"


	fmt.Printf(repeatStr, args...)
}