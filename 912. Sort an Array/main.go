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

// bubble sort
func sortArray2(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 0; i < n; i++ {
		for j := 0; j < (n - 1 - i); j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}

// selection sort
func sortArray3(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 0; i < n-1; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		temp := nums[i]
		nums[i] = nums[index]
		nums[index] = temp
	}
	return nums
}

// insertion sort
func sortArray4(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 1; i < n; i++ {
		previous := i - 1
		current := nums[i]
		for previous >= 0 && nums[previous] > current {
			nums[previous+1] = nums[previous]
			previous--
		}
		nums[previous+1] = current
	}
	return nums
}

func main() {
	nums := []int{5, 7, 1, 1, 2, 0, 0, 8, 1, 4, 6, 3, 10, 2, 9}
	fmt.Println("input Array: ", nums)
	fmt.Println("Sorted Array: ", sortArray4(nums))
}
