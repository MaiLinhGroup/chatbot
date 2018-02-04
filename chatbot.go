package main

// The chatbot handles the communication of the program with the
// underlying bot API. Therefore it should offer abstraction to
// be decoupled from the bot API. The public methods shouldn't be
// bot API specific.

import (
	"log"
)

const (
	//Telegram: @ChatLotteBot
	botAPIToken = "503887514:AAHOnl7OiyDk6oBPvHuJBEadlBOxFTnGxlk"
	botPassword = "test" // TODO: to be removed in production
)

type Update struct {
}

// ChatBot is used by the program to interact
// with users through a bot
type ChatBot struct {
	bot      *telegrambotapi // reference to telegram bot API
	password string          // TODO: password service to generate a strong password
}

// NewChatBot : creating a new chatbot instance
func (me *ChatBot) NewChatBot() {
	err := me.bot.newBotAPI(botAPIToken)
	handleBotApiError(err)
}

// GetBotUpdates : fetch update responses from bot
func (me *ChatBot) GetBotUpdates() {
}

func handleBotApiError(err error) {
	if err != nil {
		log.Fatalf("Error from telegram bot api: %v\n", err)
	}
}
