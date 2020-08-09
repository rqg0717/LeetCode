package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)

	if n2 == 0 {
		return 0
	}

	if n1 == 0 || n1 < n2 {
		return -1
	}

	for i := 0; i <= n1-n2; i++ {
		if haystack[i:i+n2] == needle {
			return i
		}
	}
	return -1
}

// Daniel M.Sunday: A Very Fast Substring Search Algorithm.
func strStr1(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n2 == 0 {
		return 0
	}
	if n1 == 0 || n1 < n2 {
		return -1
	}
	index1 := 0 // haystack
	index2 := 0 // needle
	for index2 < n2 {
		if index1 > n1-1 {
			return -1
		}
		if haystack[index1] == needle[index2] {
			index1++
			index2++
		} else {
			next := index1 - index2 + n2
			if next < n1 {
				last := strings.LastIndexByte(needle, haystack[next])
				if last == -1 {
					index1 = next + 1
				} else {
					index1 = next - last
				}
				index2 = 0
			} else {
				return -1
			}
		}
	}
	return index1 - index2
}

func main() {
	haystack := "Some books are to be tasted, others to be swallowed, and some few to be chewed and digested."
	needle := "ll"
	fmt.Println("Output: ", strStr1(haystack, needle))
}
