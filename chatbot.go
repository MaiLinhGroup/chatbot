package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

}

//DRVBot serves as personal assistant chatbot
type DRVBot struct {
	BaseBot
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
	s := string(data)

	fmt.Println(s)
}

//SetUpChatBot uses the chatbot information to set up the chat bot with
//a chatbot API (e.g. Telegram-Bot-API)
func SetUpChatBot(cb ChatBot) {

}

func main() {
	fmt.Println("chat bot main func")
}
