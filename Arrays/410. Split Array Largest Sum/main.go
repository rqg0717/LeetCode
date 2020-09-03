package main

import (
	"fmt"
)

func splitArray(nums []int, m int) int {
	left, right := 0, 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if left < nums[i] {
			left = nums[i]
		}
		right += nums[i]
	}
	for left < right {
		mid := left + (right-left)>>1
		if checkSum(nums, mid, m, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func checkSum(nums []int, mid, m, n int) bool {
	sum, counter := 0, 1
	for i := 0; i < n; i++ {
		if sum+nums[i] > mid {
			counter++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return counter <= m
}

func main() {
	nums := []int{0, 1, 2, 4, 6, 5, 1}
	m := 2
	fmt.Println("Output: ", splitArray(nums, m))
}
