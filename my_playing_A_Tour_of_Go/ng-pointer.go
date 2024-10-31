/*
	ちょっと面白いやつ。
	チャネルを使う理由は同期のため。
	goroutineはスレッドを分けて非同期処理するので、チャネルは送受信をキチンと待つ。
	無理やりポインターとスリープで実験してみると適切に待機しないと壊れることが分かる。
*/

package main

import (
	"fmt"
	"time"
)

func sum(s []int, p *int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	*p = sum // bad pointer
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	u, v := 0, 0
	x, y := &u, &v
	go sum(s[:len(s)/2], x)
	time.Sleep(1500 * time.Millisecond)
	go sum(s[len(s)/2:], y)
	// time.Sleep(1500 * time.Millisecond)
	fmt.Println(*x, *y, (*x)+(*y))
}
