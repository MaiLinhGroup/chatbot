package main

// Wrapper for the golang telegram bot api :
// github.com/go-telegram-bot-api/telegram-bot-api

// Used to abstract the golang telegram bot api from
// the own code. This code is not tested yet and will not be.

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func startTgBot(botAPIToken string) {
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
