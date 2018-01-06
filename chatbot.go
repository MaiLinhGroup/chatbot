package main

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	//@ChatLotteBot
	botAPIToken = "503887514:AAHOnl7OiyDk6oBPvHuJBEadlBOxFTnGxlk"
)

type ChatBot struct {
	bot      *tgBotAPI.BotAPI
	password string
}

func (me *ChatBot) Start() error {
	bot, err := tgBotAPI.NewBotAPI(botAPIToken)
	me.bot = bot

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgBotAPI.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgBotAPI.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	return err
}
