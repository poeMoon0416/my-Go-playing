package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

var templates = []string{"Hi! %v.", "Hello, %v.", "Good Morning %v."}

func Hello(name string) (string, error) {
	if name == "" {
		// return "error!", errors.New("input name!")
		return "", errors.New("input name")
	}
	// return fmt.Sprint(templates[rand.Intn(len(templates))]), nil
	return fmt.Sprintf(templates[rand.Intn(len(templates))], name), nil
}

/*
func Hellos(names []string) (map[string]string, error) {
	if names == nil {
		return nil, errors.New("names is <nil>")
	}
	msgs := make(map[string]string, len(names))
	for _, name := range names {
		if name == "" {
			continue
		}
		msgs[name] = fmt.Sprintf(templates[rand.Intn(len(templates))], name)
	}
	return msgs, nil
}
*/

func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string, len(names))
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}
	return msgs, nil
}
