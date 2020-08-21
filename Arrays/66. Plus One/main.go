package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carryover := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + carryover
		carryover = 0
		if i == len(digits)-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			carryover = 1
			digits[i] = digits[i] % 10
		}
	}
	if carryover > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{1, 3, 4}
	fmt.Println("Output: ", plusOne(digits))
}
