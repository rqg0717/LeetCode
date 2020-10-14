package main

import (
	"fmt"
)

func findMaximumXOR(nums []int) int {
	result := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= (1 << i)
		m := make(map[int]bool)
		for _, num := range nums {
			m[num&mask] = true
		}
		tmp := result | (1 << i)
		for j := range m {
			if _, ok := m[tmp^j]; ok {
				result = tmp
				break
			}
		}
	}
	return result
}

func main() {
	nums := []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}
	fmt.Println("Output: ", findMaximumXOR(nums))
}
