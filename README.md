Project 'Chet Bot' :gun::sunglasses: [![Build Status](https://travis-ci.org/MaiLinhGroup/chatbot.svg?branch=master)](https://travis-ci.org/MaiLinhGroup/chatbot)
==================
## Libraries
This application is using following 3rd parties

+ [tgbotapi](https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api) for interacting with the Telegram Bot API

-------------
## Dependency Management Tool
+ [dep](https://golang.github.io/dep/) : Dependency management for Go
-------------
## Prerequisites
### Telegram Bot
To interact with the Telegram Bot API, you first need to create a Bot to get an authorization token. More information and the instruction on how to create a Telegram Bot can be found [here](https://core.telegram.org/bots).
### Passwords and Authentification
This application uses environment variables heavily to set or retrieve sensible informations like keys, token, password etc. for authentification. You can set them on the fly directly in the terminal or via a bash script like this:
```bash
#!/bin/bash

CHATID="1234"
KEY="abcd"
```
Of course in both cases it's only available locally. Here are the required environment variables used by the application:

+ TOKEN : authorization token provided by the Telegram Bot API

-------------
