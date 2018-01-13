package main

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	//@ChatLotteBot
	botAPIToken = "503887514:AAHOnl7OiyDk6oBPvHuJBEadlBOxFTnGxlk"
	botPassword = "test"
)

type ChatBot struct {
	bot      *telegrambotapi
	password string
}

func (me *ChatBot) NewChatBot() {
	err := me.bot.newBotAPI(botAPIToken)
	if err != nil {
		tgbotAPIErrorHandler(err)
	}
}

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
