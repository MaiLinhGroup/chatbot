package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	//@ChatLotteBot
	botAPIToken = "503887514:AAHOnl7OiyDk6oBPvHuJBEadlBOxFTnGxlk"
)

func main() {
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
