package main

import (
	"fmt"
)

func countTriplets(arr []int) int {
	count, tmp := 0, 0
	for k := 1; k < len(arr); k++ {
		for i := 0; i < k; i++ {
			for j := i; j <= k; j++ {
				tmp ^= arr[j]
			}
			if tmp == 0 {
				count += (k - i)
			}
			tmp = 0
		}

	}
	return count
}

func countTriplets1(arr []int) int {
	m, tmp, count := map[int][]int{}, 0, 0
	m[0] = []int{-1}
	for k, val := range arr {
		tmp ^= val
		if j, ok := m[tmp]; !ok {
			m[tmp] = []int{k}
		} else {
			for _, i := range j {
				count += k - i - 1
			}
			m[tmp] = append(m[tmp], k)
		}
	}
	return count
}

func main() {
	arr := []int{2, 3, 1, 6, 7}
	fmt.Println("Output: ", countTriplets1(arr))
}
