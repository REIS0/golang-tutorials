package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set the properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing the
	// time, souce file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladus", "Dio", "Solaris"}

	// Request a greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console
	// and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned
	// messages to the console.
	fmt.Println(messages)
}
