package main

import "fmt"

// quick sort
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return []int{nums[1], nums[0]}
		}
		return nums
	}
	var low, high []int
	for _, n := range nums[1:] {
		if n < nums[0] {
			low = append(low, n)
		} else {
			high = append(high, n)
		}
	}
	var rerults = append(sortArray(low), nums[0])
	rerults = append(rerults, sortArray(high)...)
	return rerults
}

// bucket sort
func sortArray1(nums []int) []int {
	//Constraints:
	//1 <= nums.length <= 50000
	//-50000 <= nums[i] <= 50000
	bucket := make([]int, 100001)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]+50000]++
	}
	index := 0
	for i := 0; i < 100001; i++ {
		for bucket[i] > 0 {
			nums[index] = i - 50000
			index++
			bucket[i]--
		}
	}
	return nums
}

func main() {
	nums := []int{5, 7, 1, 1, 2, 0, 0, 8, 1, 4, 6, 3, 10, 2}
	fmt.Println("input Array: ", nums)
	fmt.Println("Sorted Array: ", sortArray(nums))
}
