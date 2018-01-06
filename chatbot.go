package main

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	//@ChatLotteBot
	botAPIToken = "503887514:AAHOnl7OiyDk6oBPvHuJBEadlBOxFTnGxlk"
	botPassword = "test"
)

type ChatBot struct {
	bot      *tgBotAPI.BotAPI
	password string
}

func (me *ChatBot) NewChatBot() error {
	bot, err := tgBotAPI.NewBotAPI(botAPIToken)
	me.bot = bot
	me.password = botPassword
	return err
}

func (me *ChatBot) Start() error {
	me.bot.Debug = true

	log.Printf("Authorized on account %s", me.bot.Self.UserName)

	u := tgBotAPI.NewUpdate(0)
	u.Timeout = 60

	updates, err := me.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgBotAPI.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		me.bot.Send(msg)
	}

	return err
}
