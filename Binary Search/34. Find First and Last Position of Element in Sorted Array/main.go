package main

import "fmt"

func searchFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchLast(nums, target)}
}

func findIndex(nums []int, target int, isLeft bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target || (isLeft && target == nums[mid]) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchRange1(nums []int, target int) []int {
	targetRange := []int{-1, -1}
	left := findIndex(nums, target, true)
	if left == len(nums) || nums[left] != target {
		return targetRange
	}
	targetRange[0] = left
	targetRange[1] = findIndex(nums, target, false) - 1
	return targetRange
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 9, 10, 11, 12, 12}
	target := 8
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", searchRange1(nums, target))
}
