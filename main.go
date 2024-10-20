package main

import (
	"fmt"
	"math"
)

const EPS = 0.00000000000001

func Sqrt(x float64) float64 {
	z, prev := 1.0, 0.0

	for math.Abs(z-prev) >= EPS {
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println("Closest square to 49 is", Sqrt(49))
	fmt.Println("Closest square to 50 is", Sqrt(50))
}
