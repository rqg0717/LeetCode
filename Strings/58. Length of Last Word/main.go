package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	n := len(s)
	if n == 0 {
		return 0
	}
	length := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length == 0 {
				continue
			}
			break
		}
		length++
	}
	return length
}

func main() {
	s := "  A man, a plan, a canal: Panama   "
	fmt.Println("Output: ", lengthOfLastWord(s))
}
