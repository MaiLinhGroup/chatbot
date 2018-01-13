package main

import (
	"fmt"
)

var cb *ChatBot

func main() {
	fmt.Println("main of chatbot")
	cb = &ChatBot{
		bot: &telegrambotapi{},
	}
	cb.NewChatBot()
	cb.Start()
}
