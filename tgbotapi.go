package main

// Wrapper for the golang telegram bot api :
// github.com/go-telegram-bot-api/telegram-bot-api

// Used to abstract the golang telegram bot api from
// the own code. This code is not tested yet and will not be.

import (
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
