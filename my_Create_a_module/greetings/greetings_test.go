package greetings

import (
	"regexp"
	"testing"
)

/*
func TestHelloName(t *testing.T) {
	msg, err := Hello("taro")
	matched, err := regexp.MatchString(".*taro.*", msg)
	if !matched || err != nil {
		t.Fatalf("Hello(\"taro\") return msg: %q, err: %v. But want msg: regexp\".*taro.*\", err: <nil>", msg, err)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf("Hello(\"\") return msg: %q, err: %T %q. But want msg: \"\", err: *errors.errorString \"input name!\".", msg, err, err)
	}
}
*/

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	// 正規表現として解釈できればその*RegExpが返る
	// \bは文頭、文末、単語の境界(数値、英字、アンダースコア以外の文字)を表すようだ。
	// Goにおいては非単語、単語、非単語または単語、非単語、単語が条件を満たす。
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	// RegExpをレシーバとする方のMatchString()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
