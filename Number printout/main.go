package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	output := []int{}
	if n == 0 {
		return output
	}
	length := 1
	for n > 0 {
		length = length * 10
		n--
	}
	for i := 1; i < length; i++ {
		output = append(output, i)
	}
	return output
}

func main() {
	n := 2
	fmt.Println("Output: ", printNumbers(n))
	printNumbers1(n)
}

func printNums(nums []byte) {
	flag := false
	for i := 0; i < len(nums); i++ {
		if nums[i] != '0' && !flag {
			flag = true
		}
		if flag {
			fmt.Printf("%s\n", nums[i:])
			break
		}
	}
}

func isFinished(nums []byte) bool {

	n := len(nums)
	extra := 0
	for i := n - 1; i >= 0; i-- {
		num := int(nums[i]-'0') + extra
		if i == n-1 {
			num++
		}

		if num >= 10 {
			if i == 0 {
				return true
			}
			extra = 1
			nums[i] = '0'
		} else {
			nums[i] = byte(num + '0')
			break
		}
	}
	return false
}

func printNumbers1(n int) {
	if n == 0 {
		return
	}
	nums := make([]byte, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = '0'
	}

	for !isFinished(nums) {
		printNums(nums)
	}
}
