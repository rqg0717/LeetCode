package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	x = x ^ y
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func hammingDistance2(x int, y int) int {
	return bits.OnesCount(uint(x) ^ uint(y))
}

func main() {
	x := 2
	y := 7
	fmt.Println("Output: ", hammingDistance2(x, y))
}
