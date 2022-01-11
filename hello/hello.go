package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Diss")
	fmt.Println(message)
}
