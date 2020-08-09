package main

import "fmt"

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

func strStr1(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)

	if n2 == 0 {
		return 0
	}

	if n1 == 0 || n1 < n2 {
		return -1
	}

	return -1
}

func main() {
	haystack := "Hello World YeaH!"
	needle := "YeaH"
	fmt.Println("Output: ", strStr(haystack, needle))
}
