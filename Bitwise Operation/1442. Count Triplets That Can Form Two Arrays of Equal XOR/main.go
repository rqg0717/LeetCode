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

func main() {
	arr := []int{2, 3, 1, 6, 7}
	fmt.Println("Output: ", countTriplets(arr))
}
