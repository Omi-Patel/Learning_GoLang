package main

import "fmt"

// dynamic array
func main() {

	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums)) // 0

	var arr = make([]int, 2, 5) // inbuilt function to make slice with length 2 - [0 0]

	// capacity -> maximum numbers of elements can fit -> third parameter
	fmt.Println(cap(arr)) // 2

	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)  // when the initial size is full, it doubles the size of the array
	fmt.Println(cap(arr)) // 2

	fmt.Println(arr)
}
