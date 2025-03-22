package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {

	// this is the best example of variadic function -> takes multiple arguments
	// fmt.Println(1, 2, 3, "hello")

	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	nums := []int{3, 4, 5, 6}
	fmt.Println(sum(nums...)) // just like spread operator in javascript
}
