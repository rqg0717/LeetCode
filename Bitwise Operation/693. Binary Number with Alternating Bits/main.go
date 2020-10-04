package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	for n != 0 {
		if n&1 == (n>>1)&1 {
			return false
		}
		n >>= 1
	}
	return true
}

func hasAlternatingBits1(n int) bool {
	n = n ^ (n >> 1)
	return n&(n+1) == 0
}

func main() {
	n := 5
	fmt.Println("Output: ", hasAlternatingBits1(n))
}
