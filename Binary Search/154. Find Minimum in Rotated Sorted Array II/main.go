package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left+1 < right {
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
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
	nums := []int{1, 3, 3, 3}
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", findMin(nums))
}
