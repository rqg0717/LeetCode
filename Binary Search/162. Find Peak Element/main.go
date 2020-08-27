package main

import "fmt"

func findPeakElement(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left, right int) int {
	if left == right {
		return right
	}
	mid := left + (right-left)>>1
	if nums[mid] > nums[mid+1] {
		return search(nums, left, mid)
	}
	return search(nums, mid+1, right)
}

func main() {
	nums := []int{1, 7, 6, 3, 2, 6, 4}
	fmt.Println("Output: ", findPeakElement(nums))
}
