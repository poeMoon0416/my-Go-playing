package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("hello"))
	fmt.Println(reverse.String("アジア"))
	fmt.Println(reverse.String("珈琲☕"))
	revI := reverse.Int(123)
	fmt.Printf("%v %T\n", revI, revI)
}
