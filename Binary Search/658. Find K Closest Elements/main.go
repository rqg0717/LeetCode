package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := left + (right-left)>>1
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 9, 10, 11, 12, 12}
	k := 2
	x := 7
	fmt.Println("input Array: ", nums)
	fmt.Println("Output: ", findClosestElements(nums, k, x))
}
