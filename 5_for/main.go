package main

import "fmt"

// for -> only construct available for looping in go
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for{

	// }

	// classic for loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	// range

	for i := range 3 {
		fmt.Println(i) // 0 1 2
	}
}
