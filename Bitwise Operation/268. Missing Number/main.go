package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	num := 0
	for i, val := range nums {
		num ^= val ^ i
	}
	return num ^ len(nums)
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("Output: ", missingNumber(nums))
}
