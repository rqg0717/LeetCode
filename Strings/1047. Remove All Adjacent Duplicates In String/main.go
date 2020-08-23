package main

import (
	"fmt"
)

func removeDuplicates(S string) string {
	letters := []byte{}
	for i := 0; i < len(S); i++ {
		if len(letters) == 0 {
			letters = append(letters, S[i])
			continue
		}
		if letters[len(letters)-1] == S[i] {
			letters = letters[:len(letters)-1]
		} else {
			letters = append(letters, S[i])
		}

	}
	return string(letters)
}

func removeDuplicates1(S string) string {
	r := []rune(S)
	for i := 0; i < len(r)-1; {
		if r[i] == r[i+1] {
			r = append(r[:i], r[i+2:]...)
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return string(r)
}

func main() {
	s := "cabbaca"
	fmt.Println("Output: ", removeDuplicates1(s))
}
