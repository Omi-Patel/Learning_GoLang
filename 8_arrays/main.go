package main

import "fmt"

func main() {

	// zeroed (default) values based on the type
	// int -> 0, bool -> false, string -> "" (empty string)

	// var nums [4]int

	nums := [3]int{1, 2, 3}

	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	fmt.Println(len(nums))

	// 2D arrays

	matrix := [2][2]int{{3, 4}, {5, 6}}

	fmt.Println(matrix)

	// when to user arrays

	// - fixed size, that is predicatble
	// - memory optimization
	// - constant time access
}
