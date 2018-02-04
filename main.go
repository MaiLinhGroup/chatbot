package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"fmt"
)

// pointer to the chatbot
var cb *ChatBot

func main() {
	// TODO: debug, remove chatter
	fmt.Println("main of chatbot")

	// initialising the chatbot
	cb = &ChatBot{
		bot: &telegrambotapi{},
	}
	cb.NewChatBot()
	cb.GetBotUpdates()
}
