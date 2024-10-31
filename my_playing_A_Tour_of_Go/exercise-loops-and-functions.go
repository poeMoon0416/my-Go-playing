package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var z float64 = 1
	for math.Abs(z*z - x) > math.Pow(10, -6) {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(334))
	fmt.Println(math.Sqrt(334))
}
