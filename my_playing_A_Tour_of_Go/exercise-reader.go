package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
	mr := MyReader{}
	b := make([]byte, 5)
	mr.Read(b)
	for _, v := range b {
		fmt.Println(string(v))
	}
}
