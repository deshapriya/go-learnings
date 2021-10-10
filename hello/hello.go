package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

//go mod edit -replace example.com/greetings=../greetings

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Kelum Deshapriya")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{
		"Kelum Deshapriya",
		"Darshi Egodage",
		"Sumanawathie",
	}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	for _, msg := range messages {
		fmt.Println(msg)
	}

}
