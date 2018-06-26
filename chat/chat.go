package chat

// The chat module interacts with users via the Telegram Bot API.
// It processes user requests and makes them accessible for other
// modules. It also sends back the results of other modules to the users.

import (
	"errors"
	"os"
	"strings"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

// monkey patching :monkey_face:
var osGetEnv = os.Getenv

// Bot contains the dependencies to leverage the Telegram Bot API
type Bot struct {
	API          *tgbot.BotAPI      // interaction with Telegram Bot API
	UpdateConfig tgbot.UpdateConfig // contains information about update request
}

// Message is constructed of a unique ID which is used to identify the user
// who has sent the request and the request itself
type Message struct {
	ID   int64
	Text string
}

// New authentificates a new Bot struct with the provided token at
// the Telegram Bot API and returns it ready-to-use to interact with the API
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
	botAPI.Debug = true

	// contains config information about updating user requests
	updateCfg := tgbot.NewUpdate(0)
	updateCfg.Timeout = 60

	return &Bot{API: botAPI, UpdateConfig: updateCfg}, nil
}

// Chat either receiving user requests via the update channel or
// sending results back to user using the send functionality of
// the Telegram Bot API
func (bot *Bot) Chat(userRequest, userFeedback chan Message) error {
	// Get update channel to receive messages from user
	updates, err := bot.API.GetUpdatesChan(bot.UpdateConfig)
	if err != nil {
		return err
	}

	// Clear all unprocessed updates after certain period of time
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	// Waiting for messages from user via update channel,
	// send user request via channel to handler
	// and waiting for its reply on the feedback channel to send it
	// back to user
	for {
		select {
		case upd := <-updates:
			urq := Message{
				ID:   upd.Message.Chat.ID,
				Text: upd.Message.Text,
			}
			userRequest <- urq
		case ufb := <-userFeedback:
			reply := tgbot.NewMessage(ufb.ID, ufb.Text)
			bot.API.Send(reply)
		}
	}
}

// HandleMessage ...
func HandleMessage(userRequest, userFeedback chan Message) {
	for msg := range userRequest {
		reversed := ReversedMessage(msg.Text)
		msg.Text = reversed
		userFeedback <- msg
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
