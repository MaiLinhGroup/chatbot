package main

// Wrapper for the golang telegram bot api :
// github.com/go-telegram-bot-api/telegram-bot-api

// Used to abstract the golang telegram bot api from
// the own code. This code is not tested yet and will not be.

import (
	"log"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegrambotapi struct {
	bot *tg.BotAPI
}

func (me *telegrambotapi) newBotAPI(botAPIToken string) error {
	b, err := tg.NewBotAPI(botAPIToken)
	if err != nil {
		return err
	}
	me.bot = b
	return nil
}

func (me *telegrambotapi) getUpdates() {
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := me.bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		me.bot.Send(msg)
	}
}
