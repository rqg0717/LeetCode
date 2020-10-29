package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxProduct(nums []int) int {
	dpMax, dpMin, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			dpMax, dpMin = max(nums[i], dpMax*nums[i]), min(nums[i], dpMin*nums[i])
		} else {
			dpMax, dpMin = max(nums[i], dpMin*nums[i]), min(nums[i], dpMax*nums[i])
		}
		result = max(dpMax, result)
	}
	return result
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println("Output: ", maxProduct(nums))
}
