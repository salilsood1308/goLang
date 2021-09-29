package myutils

import (
	"fmt"

	"rsc.io/quote"
)

func Greet1(name string) string {
	return fmt.Sprintf("Hello %v", name)
}

func MyQuote() string {
	return fmt.Sprintf(quote.Go())
}
