package main

import (
	"fmt"
	"sort"
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

// What if the given array is already sorted? How would you optimize your algorithm?
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return nums1[:k]
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println("Output: ", intersect2(nums1, nums2))
}
