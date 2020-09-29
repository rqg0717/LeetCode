package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		b = (a & b) << 1
		a = sum
	}
	return a
}

func getSum1(a int, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	b = a & b << 1
	a = sum
	return getSum(a, b)
}

func main() {
	a := -2
	b := 3
	fmt.Println("Output: ", getSum1(a, b))
}
