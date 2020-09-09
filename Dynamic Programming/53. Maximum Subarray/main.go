package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	ret := nums[0]
	if n == 1 {
		return ret
	}
	for i := 1; i < n; i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])
		ret = max(ret, nums[i])
	}
	return ret
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Output: ", maxSubArray(nums))
}
