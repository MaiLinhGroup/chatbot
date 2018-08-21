package main

// This is the entry point for the chat programm where all the
// packages are called to interact with each other.

import (
	"strings"

	// standard libs

	// 3rd party libs
	"fmt"
	log "github.com/goinggo/tracelog"
	"io/ioutil"
	"net/http"
	"os"
)

const TELEGRAM = "https://api.telegram.org/bot"

func main() {
	log.Start(log.LevelInfo)

	// load authentification informations
	tgBotToken := os.Getenv("TELEGRAM_TOKEN")
	// adminID := os.Getenv("ADMIN_ID")

	// https://api.telegram.org/bot<token>/METHOD_NAME
	getUpdatesRq := TELEGRAM + tgBotToken + "/getUpdates"

	// make a get request
	rs, err := http.Get(getUpdatesRq)
	// process reponse and handle err
	if err != nil {
		panic(err)
	}

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	fmt.Println(bodyString)

	defer func() {
		rs.Body.Close()
		log.Stop()
	}()
}

// ReverseMessage takes a message and returns it in reverse order.
// One or more leading and trailing whitespaces got to be removed,
// but no further modification will be performed on the original message.
func ReverseMessage(msg string) (reverseMsg string) {
	msg = strings.TrimSpace(msg)

	for i := len(msg) - 1; i >= 0; i-- {
		reverseMsg += string(msg[i])
	}

	return
}
