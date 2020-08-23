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

func main() {
	s := "abbaca"
	fmt.Println("Output: ", removeDuplicates(s))
}
