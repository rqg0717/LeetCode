package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Println("new nums: ", nums)
	return len(nums)
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Output: ", removeDuplicates(nums))
}
