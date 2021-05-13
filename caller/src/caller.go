package main

import (
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Shengzhe Chen")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)
}
