package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	// difference between two numbers (x and y) which were seen only once
	mask := 0
	for _, num := range nums {
		mask = mask ^ num
	}
	// rightmost 1-bit diff between x and y
	diff := mask & (-mask)
	x := 0
	for _, num := range nums {
		if num&diff != 0 {
			x = x ^ num
		}
	}
	y := mask ^ x
	return []int{x, y}
}

func main() {
	nums := []int{3, 1, 2, 2, 1, 5}
	fmt.Println("Output: ", singleNumber(nums))
}
