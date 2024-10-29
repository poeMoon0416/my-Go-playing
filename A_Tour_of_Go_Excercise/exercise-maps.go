package main

import (
	// "fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	// fmt.Println(words, len(words), cap(words))
	ret := make(map[string]int)
	for _, word := range words {
		ret[word]++
		// ret[word] += 2
	}
	return ret
}

func main() {
	// WordCount("I have a pen")
	wc.Test(WordCount)
}
