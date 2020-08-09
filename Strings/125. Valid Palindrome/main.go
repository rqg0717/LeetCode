package main

import "fmt"

func isPalindrome(s string) bool {
	var ss []rune
	n := len(s)
	if n <= 1 {
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

// is not 0-9 a-z A-Z?
func isNotChar(ch byte) bool {
	if ch <= 47 || (ch >= 58 && ch <= 64) ||
		(ch <= 96 && ch >= 91) || ch >= 123 {
		return true
	}
	return false
}

func isPalindrome1(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	left := 0
	right := n - 1
	for left < right {
		// skips symbols
		if isNotChar(s[left]) {
			left++
			continue
		}
		if isNotChar(s[right]) {
			right--
			continue
		}
		leftChar := s[left]
		rightChar := s[right]
		// converts uppercase to lowercase
		if 65 <= leftChar && leftChar <= 90 {
			leftChar += 32
		}
		if 65 <= rightChar && rightChar <= 90 {
			rightChar += 32
		}
		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println("Output: ", isPalindrome1(s))
}
