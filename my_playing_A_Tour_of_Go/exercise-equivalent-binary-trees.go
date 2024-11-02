package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	vs1, vs2 := make([]int, 0), make([]int, 0)
	cnt := 0
	for cnt < 20 {
		select {
		case v1 := <-ch1:
			vs1 = append(vs1, v1)
			cnt++
		case v2 := <-ch2:
			vs2 = append(vs2, v2)
			cnt++
		}
	}

	for i := 0; i < 10; i++ {
		// fmt.Println(vs1[i], vs2[i])
		if vs1[i] != vs2[i] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t := tree.New(2)
	fmt.Println(t)
	go Walk(t, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	go Walk(nil, nil)
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
