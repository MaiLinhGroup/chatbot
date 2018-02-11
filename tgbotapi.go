package main

// Wrapper for the golang telegram bot api :
// github.com/go-telegram-bot-api/telegram-bot-api

// Used to abstract the golang telegram bot api from
// the own code. This code is not tested yet and will not be.

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Update struct {
	from      string
	text      string
	messageID int
	chatID    int64
}

var bot *tgbotapi.BotAPI

func startBot(token string) {
	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot = b
}

func getUpdates() Update {
	udp := Update{}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		udp.updWrapper(update)
		return udp
	}
	return udp
}

func (me *Update) updWrapper(u tgbotapi.Update) {
	me.from = u.Message.From.UserName
	me.text = u.Message.Text
	me.messageID = u.Message.MessageID
	me.chatID = u.Message.Chat.ID
}

//

// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// msg.ReplyToMessageID = update.Message.MessageID

// bot.Send(msg)
