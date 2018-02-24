package main

import (
	"log"
	"os"
)

// The chatbot handles the communication of the program with the
// underlying bot API. Therefore it should offer abstraction to
// be decoupled from the bot API. The public methods shouldn't be
// bot API specific.

// monkey patching
var updateRetriever func() UpdatesChannel

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
	ur := UserRequest{}
	// for upd := range updateRetriever() {
	// 	log.Printf("Receiver [%s] %s", upd.from, upd.text)
	// 	ur.chatID = upd.chatID
	// 	ur.msg = upd.text
	// }
	updates := updateRetriever()
	log.Println("before select")
	select {
	case upd := <-updates:
		log.Printf("Receiver [%s] %s", upd.from, upd.text)
	default:
		log.Printf("No message received")
	}

	return ur
}

func Reply(r string) {
	send(r)
}
