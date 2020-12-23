package main

import (
	"fmt"

	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote.Go())
	fmt.Println(message)
}
