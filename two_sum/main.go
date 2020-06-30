package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	results := twoSum(nums, target)
	fmt.Println(results[0], results[1])
}
