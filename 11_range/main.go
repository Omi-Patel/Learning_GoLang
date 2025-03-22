package main

import "fmt"

// iterating over data structure
func main() {
	// nums := []int{1, 2, 3}

	// fmt.Println(nums)

	// for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// range
	// sum := 0
	// for i, num := range nums { // i -> index
	// 	// sum += num
	// 	fmt.Println(i, " -> ", num)
	// }
	// fmt.Println(sum)

	// iterating over map
	// m := map[string]string{"name": "om", "surname": "patel"}

	// for k, v := range m {
	// 	fmt.Println(k, " : ", v)
	// }

	// iterating over string

	// c -> unicode code point rune
	// i -> starting byto of rune
	for i, c := range "golang" {
		// fmt.Println(i, ":", c)
		fmt.Println(i, ":", string(c))
	}
}
