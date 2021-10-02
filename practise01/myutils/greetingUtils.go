package myutils

import (
	"fmt"

	"rsc.io/quote"
)

func Greet1(name string) string {
	return fmt.Sprintf("Hello %v", name)
}

func Greet2() string {
	return ("Hello World! Courtsey SS-Modules")
}

func Myquote() string {
	return quote.Go()
}
