package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Shengzhe Chen")
	fmt.Println(message)
}
