// 正規表現として解釈できればその*RegExpが返る
// \bは文頭、文末、単語の境界(数値、英字、アンダースコア以外の文字)を表すようだ。
// Goにおいては非単語、単語、非単語または単語、非単語、単語が条件を満たす。
// これはあまり一般的な挙動ではないようだ。
// https://pkg.go.dev/regexp/syntax@go1.23.2
// https://chatgpt.com/c/67267e23-4224-800c-acf7-de3e683344e3

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// true
	re := regexp.MustCompile(`\b_\b`)
	fmt.Println(re.MatchString(" _ "))

	// true
	re = regexp.MustCompile(`\b!\b`)
	fmt.Println(re.MatchString("a!a"))

	// false
	re = regexp.MustCompile(`\b!\b`)
	fmt.Println(re.MatchString(" ! "))
}
