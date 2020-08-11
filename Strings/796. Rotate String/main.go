package main

import (
	"fmt"
)

func rotateString(A string, B string) bool {
	n1, n2 := len(A), len(B)
	if A == B {
		return true
	}
	if n1 != n2 {
		return false
	}
	AA := A + A
	n1 = len(AA)
	for i := 0; i <= n1-n2; i++ {
		if AA[i:i+n2] == B {
			return true
		}
	}
	return false
}

func main() {
	A := ""
	B := ""
	fmt.Println("Output: ", rotateString(A, B))
}
