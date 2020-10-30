package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	sum := 0
	result := ""

	for n1 >= 0 || n2 >= 0 || sum != 0 {
		tmp1, tmp2 := 0, 0
		if n1 >= 0 {
			tmp1 = int(num1[n1] - '0')
			n1--
		}

		if n2 >= 0 {
			tmp2 = int(num2[n2] - '0')
			n2--
		}
		sum += tmp1 + tmp2
		result = strconv.Itoa(sum%10) + result
		sum = sum / 10
	}
	return result
}

func main() {
	num1 := "123456789"
	num2 := "4321"
	fmt.Println("result: ", addStrings(num1, num2))
}
