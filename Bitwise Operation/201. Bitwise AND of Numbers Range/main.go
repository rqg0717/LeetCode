package main

import (
	"fmt"
)

func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m != n {
		m >>= 1
		n >>= 1
		shift++
	}
	return n << shift
}

// from 461.
func rangeBitwiseAnd1(m int, n int) int {
	for m < n {
		n &= (n - 1)
	}
	return n
}

func main() {
	m := 5
	n := 7
	fmt.Println("Output: ", rangeBitwiseAnd1(m, n))
}
