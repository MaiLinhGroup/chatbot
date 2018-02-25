package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// TODO: debug, remove chatter
	fmt.Println("main of chatbot")

	// retrieve the token pass by env var to pass it to chatbot
	tkn := os.Getenv("TOKEN")
	if tkn == "" {
		// if no token is passing
		// then the user may just forget it,
		// but we know how to handle this error
		// (inform the user about our need)
		// that's why we won't panic here
		log.Fatal("Please passing the value for the bot api token to continue using the chatbot.")
	}

	// now we can call the chatbot with the token
}
