package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Println(sdata)

	//convert data and read line by line
	tmp := strings.Split(sdata, "\n")
	var line []string
	var lines []string
	for _, s := range tmp {
		line = strings.Fields(s)
		lines = append(lines, line[1])
	}
	fmt.Println(lines)
	fmt.Println(len(lines))

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

	fmt.Println(cfg.configs)
	fmt.Println(len(cfg.configs))

}

//SetUpChatBot uses the chatbot information to set up the chat bot with
//a chatbot API (e.g. Telegram-Bot-API)
func CreateChatBot() (*BaseBot, error) {
	err := errors.New("Must fail")
	return &BaseBot{}, err
}

func main() {
	fmt.Println("chat bot main func")
}
