package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// クロージャを使うと非再帰化できるという利点がありそう
	// structのpublic風味に対してclosureはprivate風味
	idx, prevs := 0, []int{1, 1}
	return func() int {
		ret := prevs[idx]
		// Goではlambda式は無名関数で代替(公式ブログのdeferでやってるやつ)
		prevs[idx] = func() int {
			sum := 0
			for _, v := range prevs {
				sum += v
			}
			return sum
		}()
		idx++
		idx %= len(prevs)
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
