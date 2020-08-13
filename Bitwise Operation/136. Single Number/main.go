package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}

func main() {
	nums := []int{3, 1, 2, 2, 1}
	fmt.Println("Output: ", singleNumber(nums))
}
