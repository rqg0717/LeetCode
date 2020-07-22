package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return nil
	}
	m := make(map[int]int)
	for i, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, i}
		}
		m[v] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target-v == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 3, 7, 11}
	target := 9
	results := twoSum(nums, target)
	fmt.Println(results[0], results[1])
}
