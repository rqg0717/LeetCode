package main

import "fmt"

func search(nums []int, left, right int) int {
	if left == right {
		return nums[right]
	}
	mid := left + (right-left)>>1
	if nums[left] < nums[mid] {
		return search(nums, left, mid)
	}
	return search(nums, mid+1, right)
}

func findMin(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", findMin(nums))
}
