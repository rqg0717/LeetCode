package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] > nums[right] {
		return nums[right]
	}
	return nums[left]
}

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 0, 1}
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", findMin(nums))
}
