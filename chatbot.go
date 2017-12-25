package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/robfig/cron"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	fileName = "bot-config.txt"
)

//ChatBot interface with collections of methods defined for chatbots
type ChatBot interface {
	//some basic methods every chatbot should have or offer
}

//BaseBot is base type for all chatbots
type BaseBot struct {
	//some basic informations/properties every chatbot should have
	FirstName string
	LastName  string
	UserName  string
	ID        int
	Token     string
}

//DRVBot serves as personal assistant chatbot
type DRVBot struct {
	BaseBot
	PassPhrase string
}

//ChatBotCfg is code representation of the content of the chatbot config file
type ChatBotCfg struct {
	configs map[string]BaseBot
}

//ReadChatBotCfg opens a config file with informations about the bot
//and stores the config informations
func (cfg *ChatBotCfg) ReadChatBotCfg() {
	//check if there is any config file to open
	//if not exit, otherwise proceed
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatalf("File %v does not exist. You may check spelling?\n", fileName)
	}
	log.Printf("File %v does exist.\n", fileName)

	//read the file as byte slice and convert to string
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot read %v\n.", fileName)
	}
	sdata := string(data)
	// fmt.Println(sdata)

	//convert data and read line by line
	tmp := strings.Split(sdata, "\n")
	var line []string
	var lines []string
	for _, s := range tmp {
		line = strings.Fields(s)
		lines = append(lines, line[1])
	}
	// fmt.Println(lines)
	// fmt.Println(len(lines))

	//create a map for chatbot config and instantiate the chatbot config with it
	cfg.configs = make(map[string]BaseBot)

	//instatiate config for DRVBot
	userid, _ := strconv.Atoi(lines[3])
	cfg.configs["DRVBot"] = BaseBot{FirstName: lines[0],
		LastName: lines[1],
		UserName: lines[2],
		ID:       userid,
		Token:    lines[4]}

	//instatiate config for NFPBot
	userid, _ = strconv.Atoi(lines[8])
	cfg.configs["NFPBot"] = BaseBot{FirstName: lines[5],
		LastName: lines[6],
		UserName: lines[7],
		ID:       userid,
		Token:    lines[9]}

	// fmt.Println(cfg.configs)
	// fmt.Println(len(cfg.configs))

}

//GetNewChatBot creates a chatbot with the Telegram Bot API
func GetNewChatBot(bcfg BaseBot) (*tgbotapi.BotAPI, error) {
	//create bot with Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(bcfg.Token)
	if err != nil {
		log.Panic(err)
	}
	//enable debug mode
	bot.Debug = true

	//edit bot credits
	bot.Self.FirstName = bcfg.FirstName
	bot.Self.LastName = bcfg.LastName
	bot.Self.UserName = bcfg.UserName
	bot.Self.ID = bcfg.ID

	return bot, err
}

func main() {
	fmt.Println("chat bot main func")

	//flags
	var card bool

	//cron job runner
	c := cron.New()
	c.Start()

	cbCfg := &ChatBotCfg{}
	cbCfg.ReadChatBotCfg()

	walter := cbCfg.configs["DRVBot"]
	if walter.ID != 490569313 {
		return
	}

	bot, err := GetNewChatBot(walter)
	if err != nil {
		return
	}
	log.Printf("Hello from your telegram bot %v %v!", bot.Self.FirstName, bot.Self.LastName)
	log.Printf("Authorized on account %v with id: %v", bot.Self.UserName, bot.Self.ID)

	uCfg := tgbotapi.NewUpdate(0)
	uCfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(uCfg)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		//check if user message contains a command
		switch update.Message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello "+update.Message.From.UserName+", my name is "+bot.Self.FirstName+" "+bot.Self.LastName+"!\nHow can I help you?")
			bot.Send(msg)
		case "entry":
			if card {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry "+update.Message.From.UserName+", you already started this job.")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				continue
			}
			card = true
			//Fire at 07:15 AM every Monday, Tuesday, Wednesday, Thursday and Friday
			c.AddFunc("0 15 7 ? * MON-FRI", func() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hey "+update.Message.From.FirstName+", don't forget your entry cards and have a nice day honey bun!")
				bot.Send(msg)
			})
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Start cron job: entry")
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry "+update.Message.From.UserName+", I did not understand you.")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}
