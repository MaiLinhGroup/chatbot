package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/robfig/cron"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	fileName  = "bot-config.txt"
	denialTxt = "Sorry, we should get to know each other better before I can do this with you."
)

var chatID int64
var userName string

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

func chat(bot *tgbotapi.BotAPI, chatMsg string) {
	msg := tgbotapi.NewMessage(chatID, chatMsg)
	bot.Send(msg)
}

func interactionWithUser(bot *tgbotapi.BotAPI) {
	//cron job runner
	c := cron.New()
	c.Start()

	uCfg := tgbotapi.NewUpdate(0)
	uCfg.Timeout = 60

	updates, _ := bot.GetUpdatesChan(uCfg)

	for update := range updates {
		//debug purpose
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//no message: do nothing and wait for user interaction
		if update.Message == nil {
			continue
		}

		//check if user message contains a command
		if !update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, I didn't understand you.")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			continue
		}

		//commands
		chatID = update.Message.Chat.ID
		userName = update.Message.From.UserName
		if userName == "" {
			userName = update.Message.From.FirstName
		}
		switch update.Message.Command() {
		case "start":
			chat(bot, "Hello "+userName+"! My name is "+bot.Self.FirstName+" "+bot.Self.LastName+" and I'm your telegram bot.\nNow we know each other, how can I help you?")
		case "help":
			chat(bot, "You can get a list of available commands by type in /commands.\nSome commands need further arguments, and some additionally need you to authenticate yourself!")
		case "reminder":
			args := update.Message.CommandArguments()

			if args == "" {
				chat(bot, "You forgot to tell me when and what I should remind you to do.\nI need a valid cron expression followed by a colon and then the reason for the reminder.\nLet's have another try ;)!")
				continue
			}

			argsS := strings.Split(args, ":")

			if len(argsS) < 2 {
				chat(bot, "Not enough arguments for the reminder. I need a valid cron expression followed by a colon and then the reason for the reminder.\nLet's have another try ;)!")
				continue
			}

			cExpr := argsS[0]
			cause := argsS[1]

			//default cause
			if cause == "" {
				err := c.AddFunc(cExpr, func() {
					chat(bot, "Hey "+userName+", you're not my supervisor!\nWell I'd remind you to do something,but I forgot what it was, so whatever, just do it :P!")
				})
				if err != nil {
					chat(bot, "Reminder job cancelled because something went wrong with the cron expression "+cExpr)
					continue
				}

				chat(bot, "Ok, cron job "+cExpr+" is added.")
				continue
			}

			//specific cause
			err := c.AddFunc(cExpr, func() {
				chat(bot, "Hi "+userName+", here is "+bot.Self.FirstName+". Don't forget to do this: \n"+cause)
			})
			if err != nil {
				chat(bot, "Reminder job cancelled because something went wrong with the cron expression "+cExpr)
				continue
			}
			chat(bot, "Ok, cron job "+cExpr+" is added and the reminder is:\n"+cause)
		case "stop":
			args := update.Message.CommandArguments()
			if args != "reminder" {
				c.Stop()
				chat(bot, "This command needs an argument, please try again.")
				continue
			}
			chat(bot, "Stop all available reminder cron jobs.")
		case "restart":
			args := update.Message.CommandArguments()
			if args != "reminder" {
				c.Start()
				chat(bot, "This command needs an argument, please try again.")
				continue
			}
			chat(bot, "Restart all available reminder cron jobs.")
		case "delete":
			args := update.Message.CommandArguments()
			if args == "reminder" && update.Message.From.UserName == "MLEdith" {
				c.Stop()
				c = nil
				c = cron.New()
				chat(bot, update.Message.From.UserName+": Delete all reminder cron jobs.")
				continue
			}
			chat(bot, denialTxt)
		case "new":
			args := update.Message.CommandArguments()
			if args == "reminder" && update.Message.From.UserName == "MLEdith" {
				c.Start()
				chat(bot, update.Message.From.UserName+": New cron job runner.")
				continue
			}
			chat(bot, denialTxt)
		default:
			chat(bot, "Sorry the command "+update.Message.Text+" is not available yet or unknown.")
		}
	}
}

//TravelTime calculates the travel time starting from leaving the house door and entering the workplace and vise versa.
//The rules for calculation is:
//
func TravelTime(t ...string) (duration string) {
	duration = ""
	return
}

func main() {
	cbCfg := &ChatBotCfg{}
	cbCfg.ReadChatBotCfg()

	walter := cbCfg.configs["DRVBot"]
	if walter.ID != 490569313 {
		log.Println("Missmatch, please check the bot ID.")
		return
	}

	bot, err := GetNewChatBot(walter)
	if err != nil {
		return
	}

	log.Printf("Hello from your telegram bot %v %v!", bot.Self.FirstName, bot.Self.LastName)
	log.Printf("Authorized on account %v with id: %v", bot.Self.UserName, bot.Self.ID)

	interactionWithUser(bot)
}
