package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// #3
func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[left] >= target {
		return left
	} else if nums[right] >= target {
		return right
	} else if nums[right] < target {
		return right + 1
	}
	return 0
}

func main() {
	nums := []int{1, 3, 5, 6, 7, 9}
	target := 2
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", searchInsert1(nums, target))
}
