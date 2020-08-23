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

func main() {
	nums := []int{5, 1, 3, 4, 2, 3}
	fmt.Println("Output: ", findDuplicate(nums))
}
