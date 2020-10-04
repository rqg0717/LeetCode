package main

import (
	"fmt"
)

func findComplement(num int) int {
	result := 1
	for result <= num {
		result <<= 1
	}
	result = (result - 1) ^ num
	return result
}

func main() {
	num := 5
	fmt.Println("Output: ", findComplement(num))
}
