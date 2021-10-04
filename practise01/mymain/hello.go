package main

import (
	"fmt"
	"log"

	myutils "github.com/salilsood1308/goLang/practise01/myutils"
)

func main() {
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	fmt.Println(myutils.Greet2())
	fmt.Println(myutils.Myquote())

	message, err := myutils.Greet1("Ishan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	mesg, err := myutils.Greet1("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mesg)
}
