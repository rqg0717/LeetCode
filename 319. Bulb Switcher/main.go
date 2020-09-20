package main

import "fmt"

func bulbSwitch(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count%2 != 0 { // on when odd
			result++
		}
	}
	return result
}

func main() {
	n := 16
	fmt.Println("Output: ", bulbSwitch(n))
}
