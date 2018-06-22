package chat

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

// monkey patching :monkey_face:
var osGetEnv = os.Getenv

// Bot ...
type Bot struct {
	API          *tgbot.BotAPI      // interaction with Telegram Bot API
	UpdateConfig tgbot.UpdateConfig // contains information about update request
}

// Message ...
type Message struct {
	ID   int64
	Text string
}

// New ...
func New() (*Bot, error) {
	// Each bot is given a unique authentication token when it is created
	token := osGetEnv("TOKEN")
	if token == "" {
		return nil, errors.New("bot token's missing")
	}

	// call the chatbot with the provided token
	botAPI, err := tgbot.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// uncomment this for tgbot debug
	// botAPI.Debug = true

	updateCfg := tgbot.NewUpdate(0)
	updateCfg.Timeout = 60

	return &Bot{API: botAPI, UpdateConfig: updateCfg}, nil
}

// Chat ...
func (bot *Bot) Chat(msgCh chan Message) error {
	// Get update channel to receive messages from user
	updCh, err := bot.API.GetUpdatesChan(bot.UpdateConfig)
	if err != nil {
		return err
	}

	// Clear all unprocessed updates after certain period of time
	time.Sleep(time.Millisecond * 500)
	updCh.Clear()

	for {
		select {
		case upd := <-updCh:
			msgFromUser := Message{
				ID:   upd.Message.Chat.ID,
				Text: upd.Message.Text,
			}
			fmt.Println("Trying to send msg to handler...")
			msgCh <- msgFromUser
			fmt.Println("Message sent to handler.")
		case msgToUser := <-msgCh:
			reply := tgbot.NewMessage(msgToUser.ID, msgToUser.Text)
			bot.API.Send(reply)
			fmt.Printf("Reply with ChatID %v and Text '%s'.\n", reply.ChatID, reply.Text)
		}
	}
}

// HandleMessage ...
func HandleMessage(msgCh chan Message) {
	fmt.Println("Waiting for messages to process...")
	for msg := range msgCh {
		fmt.Println("Message from user received.")
		reversed := ReversedMessage(msg.Text)
		msg.Text = reversed
		fmt.Println("Trying to send processed msg back...")
		msgCh <- msg
		fmt.Println("Message sent back.")
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
