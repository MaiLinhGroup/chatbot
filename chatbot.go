package main

import (
	"log"
	"os"
)

// The chatbot handles the communication of the program with the
// underlying bot API. Therefore it should offer abstraction to
// be decoupled from the bot API. The public methods shouldn't be
// bot API specific.

var updateRetriever func() Update

func init() {
	updateRetriever = getUpdates
}

type UserRequest struct {
	chatID int64
	msg    string
	cmds   map[string]string //key is cmd and value is arg
}

func StartChatBot() {
	tkn := os.Getenv("TKN")
	if tkn == "" {
		log.Fatal("No API token for telegram bot found.")
	}
	startBot(tkn)
}

func HandleUserRequest() UserRequest {
	udp := updateRetriever()
	log.Printf("[%s] %s", udp.from, udp.text)
	return UserRequest{
		chatID: udp.chatID,
		msg:    udp.text,
	}
}

func Reply(r string) {
	send(r)
}
