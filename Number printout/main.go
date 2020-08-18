package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	output := []int{}
	length := 0
	for n > 0 {
		length = length*10 + 9
		n--
	}
	for i := 0; i < length+1; i++ {
		output = append(output, i)
	}
	return output
}

func main() {
	n := 4
	fmt.Println("Output: ", printNumbers(n))
}
