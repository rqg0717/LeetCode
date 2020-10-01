package main

import "fmt"

func bin(num int) int {
	result := 0
	for num > 0 {
		if num%2 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}

func readBinaryWatch(num int) []string {
	results := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j <= 59; j++ {
			if bin(i)+bin(j) == num {
				results = append(results, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return results
}

func main() {
	num := 1
	fmt.Println("Output: ", readBinaryWatch(num))
}
