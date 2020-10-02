package main

import "fmt"

func toHex(num int) string {
	if num < 0 {
		num += 0x100000000
	}
	ch := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	result := []byte{}
	for {
		temp := num % 16
		result = append([]byte{ch[temp]}, result...)
		num >>= 4
		if num == 0 {
			break
		}
	}
	return string(result)
}

func main() {
	num := -4294967296
	fmt.Println("Output: ", toHex(num))
}
