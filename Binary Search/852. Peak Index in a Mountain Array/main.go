package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	fmt.Println("Output: ", peakIndexInMountainArray(nums))
}
