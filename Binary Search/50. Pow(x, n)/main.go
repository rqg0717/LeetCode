package main

import "fmt"

func mul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := mul(x, n/2)
	// if n is even, x^n = y^2
	if n%2 == 0 {
		return y * y
	}
	// if n is odd, x^n = y^2*x
	return y * y * x
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return mul(x, n)
	}
	return 1.0 / mul(x, -n)
}

func main() {
	x := 2.00000
	n := 10
	fmt.Println("Output: ", myPow(x, n))
}
