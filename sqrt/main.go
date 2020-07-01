package main

import (
	"fmt"
	"math"
)

const n = 10

// Newton's method
func Sqrt(x float64) float64 {
	z := float64(x)
	zz := float64(0)
	i := 1
	for {
		if i > n {
			return z
		}
		// adjust z based on how close z² is to x, producing a better guess
		z -= (z*z - x) / (2 * z)
		fmt.Println("i =", i, z)

		if zz == z {
			return z
		}

		zz = z
		i++
	}
}

func main() {
	x := 5.0
	fmt.Println("final result:", Sqrt(x))
	fmt.Println("math.Sqrt = ", math.Sqrt(x))
}
