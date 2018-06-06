package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"fmt"
	"os"
	"time"

	log "github.com/goinggo/tracelog"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
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
	udp.Clear()

	// TODO now echoing the sent text is working, start to do the same but with channels and go routines
	// ch := make(chan string)

	// go func() {
	// defer close(ch)

	for u := range udp {
		if u.Message == nil {
			continue
		}

		fmt.Printf("From [%s] : %s\n", u.Message.From.UserName, u.Message.Text)
		msg := tgbot.NewMessage(u.Message.Chat.ID, u.Message.Text)
		msg.ReplyToMessageID = u.Message.MessageID

		reply, err := bot.Send(msg)
		if err != nil {
			log.Error(err, "main", "bot.Send")
		}

		fmt.Printf("Reply with ID %v and Message %s.\n", msg.ReplyToMessageID, reply.Text)

		// ch <- u.Message.Text
	}
	// }()

	// for msg := range ch {
	// 	// use your string :)
	// }
}
