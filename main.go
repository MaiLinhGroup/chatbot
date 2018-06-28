package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"strings"

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

	go ChatHandler(userRq, userFb)

	chatbot.Chat(userRq, userFb)
}

// ChatHandler ...
func ChatHandler(userRequest, userFeedback chan chat.Message) {
	for msg := range userRequest {
		msg.Reply = ProcessingUserRequest(msg.Request)
		userFeedback <- msg
	}

}

// ReversedMessage takes a message and returns it in reversed order.
// One or more leading and trailing whitespaces got to be removed,
// but no further modification will be performed on the original message.
func ReversedMessage(msg string) (reversedMsg string) {
	msg = strings.TrimSpace(msg)

	for i := len(msg) - 1; i >= 0; i-- {
		reversedMsg += string(msg[i])
	}

	return
}

// ProcessingUserRequest ...
func ProcessingUserRequest(request map[string]string) string {
	return ReversedMessage(request[""])
}
