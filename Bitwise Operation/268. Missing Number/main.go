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

func missingNumber1(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	result := 0
	for _, val := range nums {
		result += val
	}
	return sum - result
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("Output: ", missingNumber1(nums))
}
