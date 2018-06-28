package chat

// The chat module interacts with users via the Telegram Bot API.
// It processes user requests and makes them accessible for other
// modules. It also sends back the results of other modules to the users.

import (
	"errors"
	"os"
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

// Message is constructed of a unique ID which is used to identify the chat/source
// where the user request has come from, information about the user who has sent the
// request and the request itself
type Message struct {
	ID      int64
	From    User
	Request map[string]string
}

// User ...
type User struct {
	ID       int
	UserName string
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
			u := User{upd.Message.From.ID, upd.Message.From.UserName}

			rq := make(map[string]string)
			if upd.Message.IsCommand() {
				rq[upd.Message.Command()] = upd.Message.CommandArguments()
			} else {
				rq[""] = upd.Message.Text
			}

			urq := Message{
				ID:      upd.Message.Chat.ID,
				From:    u,
				Request: rq,
			}
			userRequest <- urq
		case ufb := <-userFeedback:
			reply := tgbot.NewMessage(ufb.ID, ufb.Request[""])
			bot.API.Send(reply)
		}
	}
}
