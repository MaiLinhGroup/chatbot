package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	// standard libs
	"fmt"
	"os"
	"strings"
	"time"

	// 3rd party libs
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/goinggo/tracelog"
)

func main() {
	log.Start(log.LevelInfo)
	defer log.Stop()

	// retrieve the token pass by environment variable
	tkn := os.Getenv("TOKEN")
	if tkn == "" {
		log.Info("main", "os.Getenv", "TOKEN not found")
		return
	}

	// call the chatbot with the provided token
	bot, err := tgbot.NewBotAPI(tkn)
	if err != nil {
		log.Error(err, "main", "tgbot.NewBotAPI")
		return
	}

	bot.Debug = true

	ucfg := tgbot.NewUpdate(0)
	ucfg.Timeout = 60

	udp, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Error(err, "main", "bot.GetUpdatesChan")
		return
	}
	time.Sleep(time.Millisecond * 500)
	// clear all unprocessed updates after 500ms
	udp.Clear()

	ch := make(chan *tgbot.Message)

	go func() {
		defer close(ch)

		for u := range udp {
			if u.Message == nil {
				continue
			}

			fmt.Printf("From [%s] : %s\n", u.Message.From.UserName, u.Message.Text)

			ch <- u.Message
		}
	}()

	for msg := range ch {
		// use your string :)
		reversed := ReversedMessage(msg.Text)
		reply := tgbot.NewMessage(msg.Chat.ID, reversed)
		reply.ReplyToMessageID = msg.MessageID

		bot.Send(reply)

		fmt.Printf("Reply with ID %v and Message '%s'.\n", reply.ReplyToMessageID, reply.Text)
	}
}

// ReversedMessage takes a message and returns it in reversed order.
// One or more leading and trailing whitespaces got to be removed,
// but no further modification will be performed on the original message.
func ReversedMessage(msg string) (reversedMsg string) {
	msg = strings.TrimSpace(msg)

	for i := len(msg) - 1; i >= 0; i-- {
		reversedMsg += string(msg[i])
	}

	return
}
