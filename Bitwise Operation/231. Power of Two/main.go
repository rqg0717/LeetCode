package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	n := 1
	fmt.Println("Output: ", isPowerOfTwo(n))
}
