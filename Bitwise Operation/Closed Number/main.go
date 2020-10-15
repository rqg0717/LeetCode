package main

import (
	"fmt"
)

func findClosedNumbers(num int) []int {
	results := []int{}
	if num == 1 {
		results = append(results, 2, -1)
	}
	tmp := num & -num
	tmp1 := num + tmp
	tmp = ((num & ^tmp1) / tmp >> 1) | tmp1
	results = append(results, tmp)
	tmp = ^num
	tmp = tmp & -tmp
	tmp1 = ^num + tmp
	tmp = ^(((^num & ^tmp1) / tmp >> 1) | tmp1)
	results = append(results, tmp)
	return results
}

func main() {
	num := 2
	fmt.Println("Output: ", findClosedNumbers(num))
}
