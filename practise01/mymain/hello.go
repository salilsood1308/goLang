package main

import (
	"fmt"
	"log"

	myutils "github.com/salilsood1308/goLang/practise01/myutils"
)

func main() {
	log.SetPrefix("greetings: ")
	fmt.Println(myutils.Greet1("Salil"))
	fmt.Println(myutils.Greet2())
	fmt.Println(myutils.Myquote())
}
