package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func main() {
	num := uint32(0b11111111111111111111111111111101)
	fmt.Println("Output: ", hammingWeight(num))
}
