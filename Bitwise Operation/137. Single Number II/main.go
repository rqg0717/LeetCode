package main

import (
	"fmt"
)

// use the same strategy in two sum
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++ // count
	}
	for key, val := range m {
		if val == 1 {
			return key
		}
	}
	return 0
}

func singleNumber1(nums []int) int {
	a, b, c := 0, 0, 0
	for _, num := range nums {
		a = (b & ^num) | (c & num)
		c = (^b & ^c & num) | (c & ^num)
		b = a
	}
	return c
}

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println("Output:", singleNumber1(nums))
}
