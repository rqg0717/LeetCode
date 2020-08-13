package main

import (
	"fmt"
)

// use the same strategy in two sum
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++ // count
	}
	for key, val := range m {
		if val == 1 {
			return key
		}
	}
	return 0
}

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println("Output: ", singleNumber(nums))
}
