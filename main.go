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

	msgCh := make(chan chat.Message)

	go chat.HandleMessage(msgCh)

	chatbot.Chat(msgCh)
}
