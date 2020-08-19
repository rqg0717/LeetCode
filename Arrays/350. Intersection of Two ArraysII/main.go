package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	results := []int{}
	for _, value := range nums1 {
		m[value]++
	}
	for _, value := range nums2 {
		if m[value] > 0 {
			m[value]--
			results = append(results, value)
		}
	}
	return results
}

func intersect1(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, value := range nums1 {
		m[value]++
	}
	index := 0
	for _, value := range nums2 {
		if m[value] > 0 {
			m[value]--
			nums2[index] = value
			index++
		}
	}
	return nums2[0:index]
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println("Output: ", intersect(nums1, nums2))
}
