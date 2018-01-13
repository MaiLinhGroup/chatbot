package main

// TODO : move all the direct interaction with api to the tgbotapi.go
// and create wrapper methods

// The chatbot handles the communication of the program with the
// underlying bot API. Therefore it should offer abstraction to
// be decoupled from the bot API. The public methods shouldn't be
// bot API specific.

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	//Telegram: @ChatLotteBot
	botAPIToken = "503887514:AAHOnl7OiyDk6oBPvHuJBEadlBOxFTnGxlk"
	botPassword = "test" // TODO: to be removed in production
)

// ChatBot is used by the program to interact
// with users through a bot
type ChatBot struct {
	bot      *telegrambotapi // reference to telegram bot API
	password string          // TODO: password service to generate a strong password
}

// NewChatBot : method for creating a new chatbot instance
func (me *ChatBot) NewChatBot() {
	err := me.bot.newBotAPI(botAPIToken)
	if err != nil {
		tgbotAPIErrorHandler(err)
	}
}

// Start : method to start a conversation with the bot
func (me *ChatBot) Start() {
	me.bot.bot.Debug = true

	log.Printf("Authorized on account %s", me.bot.bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := me.bot.bot.GetUpdatesChan(u)
	if err != nil {
		tgbotAPIErrorHandler(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		me.bot.bot.Send(msg)
	}
}

func tgbotAPIErrorHandler(e error) {
	log.Fatalf("Something went wrong while calling the telegram bot API:\n%v", e)
}
