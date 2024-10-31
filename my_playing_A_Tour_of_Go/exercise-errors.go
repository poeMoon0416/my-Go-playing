package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// float64(e)しないとfmtのeの文字列表現探すところで無限ループで137(メモリ枯渇)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for math.Abs(z*z-x) > math.Pow(10, -6) {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	// s := ErrNegativeSqrt(-33).Error()
	// fmt.Printf("%v %T\n", s, s)
}
