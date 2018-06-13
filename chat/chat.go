package chat

import (
	"errors"
	"os"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

// monkey patching :monkey_face:
var osGetEnv = os.Getenv

// Bot ...
type Bot struct {
	botAPI       *tgbot.BotAPI      // interaction with Telegram Bot API
	updateConfig tgbot.UpdateConfig // contains information about update request
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
	botAPI.Debug = true

	updateCfg := tgbot.NewUpdate(0)
	updateCfg.Timeout = 60

	return &Bot{botAPI: botAPI, updateConfig: updateCfg}, nil
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
