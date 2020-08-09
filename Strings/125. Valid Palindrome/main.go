package main

import "fmt"

func isPalindrome(s string) bool {
	var ss []rune
	n := len(s)
	if n == 0 {
		return true
	}
	// converts all characters to lower cases and removes symbols
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
			ss = append(ss, ch)
		} else if ch >= 'A' && ch <= 'Z' {
			ss = append(ss, ch+('a'-'A'))
		} else {
			continue
		}
	}
	n = len(ss)
	for i := 0; i < n/2; i++ {
		if ss[i] != ss[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println("Output: ", isPalindrome(s))
}
