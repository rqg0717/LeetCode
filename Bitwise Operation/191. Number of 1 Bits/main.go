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

func hammingWeight1(num uint32) int {
	count := 0
	mask := uint32(1)
	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			count++
		}
		mask = mask << 1
	}
	return count
}

func main() {
	num := uint32(0b11111111111111111111111111111101)
	fmt.Println("Output: ", hammingWeight1(num))
}
