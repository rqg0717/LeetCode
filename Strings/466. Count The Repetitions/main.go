package main

import "fmt"

// Ra = ["abaacdbac", 100] then Sa = "abaacdbac", Na = 100
// Rb = ["adcbd", 4] then Sb = "adcbd", Nb = 4
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // index of Ra and Rb
	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}
	map1, map2 := make(map[int]int), make(map[int]int)
	Rb := 0                // Rb in Ra
	for index1/len1 < n1 { // Ra traversal
		if index1%len1 == len1-1 { //end of Sa
			if val, ok := map1[index2%len2]; ok { // isRepeating?
				// repeat
				repeatLen := index1/len1 - val/len1               // in Sa
				repeatCount := (n1 - 1 - index1/len1) / repeatLen // how many repeats
				SbCount := index2/len2 - map2[index2%len2]/len2   // how many Sb in a repeat

				index1 += repeatLen * repeatCount * len1 // move Ra
				Rb += repeatCount * SbCount
			} else { // first time
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}

		}
		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				Rb++
			}
			index2++
		}
		index1++
	}
	return Rb / n2
}

func main() {
	s1 := "abaacdbac"
	n1 := 100
	s2 := "adcbd"
	n2 := 4
	fmt.Println("Output: ", getMaxRepetitions(s1, n1, s2, n2))
}
