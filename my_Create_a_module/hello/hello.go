package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// log.Fatal()にはstringでなくerrorを渡す必要があり、直接エラーメッセージを編集できないため。
	log.SetPrefix("greetings: ")
	// log.Fatal()はerrorの前に"yyyy/mm/dd hh:mm:ss "をつけるため。
	// フラグはfile:lineビットと.ssssssビットとhh:mm:ssビットとyyyy/mm/ddビットで構成される。
	// log.SetFlags(31)
	// log.SetFlags(32)
	log.SetFlags(0)

	/*
		msg, err := greetings.Hello("poeMoon")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)

		msg, err = greetings.Hello("")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	*/
	/*
		msgs, err := greetings.Hellos([]string{"一郎", "次郎", "", "四郎"})
		if err != nil {
			log.Fatal(err)
		} else {
			for _, msg := range msgs {
				fmt.Println(msg)
			}
		}

		msgs, err = greetings.Hellos(nil)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, msg := range msgs {
				fmt.Println(msg)
			}
		}
	*/
	msgs, err := greetings.Hellos(nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msgs)
	}

	msgs, err = greetings.Hellos([]string{"一郎", "次郎", "三郎"})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msgs)
	}

	msgs, err = greetings.Hellos([]string{"一郎", "", "三郎"})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msgs)
	}
}
