package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// TODO: debug, remove chatter
	fmt.Println("main of chatbot")

	// retrieve the token pass by env var to pass it to chatbot
	tkn := os.Getenv("TOKEN")
	if tkn == "" {
		log.Println("Please passing a valid token.We want something like 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11.")
	}

	// now we can call the chatbot with the token
	bot, err := tgbot.NewBotAPI(tkn)
	if err != nil {
		log.Panic("in tgbot NewBotAPI:", err)
	}

	bot.Debug = true

	ucfg := tgbot.NewUpdate(0)
	ucfg.Timeout = 60

	udp, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Panic("in tgbot GetUpdatesChan:", err)
	}
	time.Sleep(time.Millisecond * 500)
	udp.Clear()

	// ch := make(chan string)

	// go func() {
	// defer close(ch)

	for u := range udp {
		if u.Message == nil {
			continue
		}

		log.Printf("From [%s] : %s\n", u.Message.From.UserName, u.Message.Text)
		msg := tgbot.NewMessage(u.Message.Chat.ID, u.Message.Text)
		msg.ReplyToMessageID = u.Message.MessageID

		reply, err := bot.Send(msg)
		if err != nil {
			log.Panic("in tgbot Send", err)
		}

		log.Printf("Reply with ID %v and %s.\n", msg.ReplyToMessageID, reply.Text)

		// ch <- u.Message.Text
	}
	// }()

	// for msg := range ch {
	// 	// use your string :)
	// }
}
