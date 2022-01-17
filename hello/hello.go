package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set props of predefined Logger, including log entry prefix
	// and a flag to disable printing the time, source file, and line no
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"General Kenobi", "Anakin", "Melon Lord", "Bill"}
	messages, err := greetings.Hellos(names)
	// if error returned, print to console and exit program
	if err != nil {
		log.Fatal(err)
	}
	// no error found
	fmt.Println(messages)
}
