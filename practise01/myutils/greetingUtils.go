package myutils

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

func Greet1(name string) (string, error) {
	// if name == "" {
	// 	return "", errors.New("Empty name")
	// }
	// message := fmt.Sprintf("Hello! %v", name)
	// return message, nil
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Greet2() string {
	return "Hello World! Courtsey SS-Modules"
}

func Myquote() string {
	return quote.Go()
}
