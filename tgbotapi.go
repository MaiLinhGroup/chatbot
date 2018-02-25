package main

// Wrapper for the golang telegram bot api :
// github.com/go-telegram-bot-api/telegram-bot-api

// Used to abstract the golang telegram bot api from
// the own code. This code is not tested yet and will not be.

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// 9 of 10 pointer semantics
type Update struct {
	from      string
	text      string
	messageID int
	chatID    int64
}

// built-in and reference types should be used with value semantics
type UpdatesChannel <-chan Update

var bot *tgbotapi.BotAPI

// don't mixing semantics, if something is using pointer semantics, stick to it!
// don't switch between value and pointer semantics!
func startBot(token string) {
	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot = b
}

//here i need to fetch updates from chat and share them with the rest of the program
func getUpdates() UpdatesChannel {
	var upd Update
	ch := make(chan Update, 100)
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
		upd.updWrapper(update)
		ch <- upd
		log.Printf("Sender [%s] %s", upd.from, upd.text)
	}

	return ch
}

func (me *Update) updWrapper(u tgbotapi.Update) {
	me.from = u.Message.From.UserName
	me.text = u.Message.Text
	me.messageID = u.Message.MessageID
	me.chatID = u.Message.Chat.ID
}

func send(m string) {
	msg := tgbotapi.NewMessage(480821480, m) //chatid where i want to send my msg to
	bot.Send(msg)
}

//https://play.golang.org/p/xrQEnuqlPXt

// msg.ReplyToMessageID = update.Message.MessageID
