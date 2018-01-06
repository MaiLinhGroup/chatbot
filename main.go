package main

import (
	"fmt"
)

var cb *ChatBot

func main() {
	fmt.Println("main of chatbot")
	cb = &ChatBot{}
	cb.NewChatBot()
	cb.Start()
}
