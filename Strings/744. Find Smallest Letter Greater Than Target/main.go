package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)>>1
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left%len(letters)]
}

func nextGreatestLetter1(letters []byte, target byte) byte {
	seen := [26]bool{}
	for _, ch := range letters {
		seen[ch-'a'] = true
	}
	for {
		target++
		if target > 'z' {
			target = 'a'
		}
		if seen[target-'a'] {
			return target
		}
	}
}

func main() {
	s := "cfj"
	letters := []byte(s)
	target := byte('a')
	fmt.Println("Output: ", nextGreatestLetter1(letters, target))
}
