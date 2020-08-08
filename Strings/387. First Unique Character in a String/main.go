package main

import "fmt"

func firstUniqChar(s string) int {
	m := make(map[string]int)
	// counts number of appearances
	for _, val := range []rune(s) {
		if _, ok := m[string(val)]; ok {
			m[string(val)]++
		} else {
			m[string(val)] = 1
		}
	}
	// finds the index of the first non-repeating letter
	for i, v := range []rune(s) {
		count, _ := m[string(v)]
		if count == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "Hello World YeaH!"
	fmt.Println("Output: ", firstUniqChar(s))
}
