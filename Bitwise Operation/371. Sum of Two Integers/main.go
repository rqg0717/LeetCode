package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func main() {
	a := -2
	b := 3
	fmt.Println("Output: ", getSum(a, b))
}
