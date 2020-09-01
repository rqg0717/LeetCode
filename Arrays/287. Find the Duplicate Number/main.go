package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = i
	}
	return -1
}

// similar to 142
func findDuplicate1(nums []int) int {
	slow, fast := 0, 0
	for {
		// slow moves one step while fast moves two steps at a time
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	fast = 0
	for {
		fast = nums[fast]
		slow = nums[slow]
		if slow == fast {
			break
		}
	}
	return fast
}

func findDuplicate2(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	result := 0
	for left <= right {
		mid := (left + right) >> 1
		counter := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				counter++
			}
		}
		if counter <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			result = mid
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println("Output: ", findDuplicate2(nums))
}
