package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // left side is ordered
			if target > nums[mid] {
				left = mid + 1
			} else if target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right side is ordered
			if target < nums[mid] {
				right = mid - 1
			} else if target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{4, 6, 7, 0, 1, 2}
	target := 5
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", search(nums, target))
}
