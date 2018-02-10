package main

import (
	"fmt"
	"os"
)

// The chatbot handles the communication of the program with the
// underlying bot API. Therefore it should offer abstraction to
// be decoupled from the bot API. The public methods shouldn't be
// bot API specific.

func StartChat() {
	fmt.Println("TOKEN:", os.Getenv("TKN"))
	startTgBot(os.Getenv("TKN"))
}
