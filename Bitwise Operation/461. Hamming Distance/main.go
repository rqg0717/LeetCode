package main

import "fmt"

func hammingDistance(x int, y int) int {
	x = x ^ y
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func main() {
	x := 2
	y := 7
	fmt.Println("Output: ", hammingDistance(x, y))
}
