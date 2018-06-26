package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"github.com/MaiLinhGroup/chatbot/chat"
	// standard libs

	// 3rd party libs
	log "github.com/goinggo/tracelog"
)

func main() {
	log.Start(log.LevelInfo)
	defer log.Stop()

	chatbot, err := chat.New()
	if err != nil {
		log.Error(err, "main", "chat.New()")
		return
	}

	userRq := make(chan chat.Message)
	userFb := make(chan chat.Message)

	go chat.HandleMessage(userRq, userFb)

	chatbot.Chat(userRq, userFb)
}
