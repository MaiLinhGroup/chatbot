package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadChatBotCfg(t *testing.T) {
	//Given
	m := make(map[string]BaseBot)
	cfg := ChatBotCfg{m}

	//When
	cfg.ReadChatBotCfg()

	//Then
	assert.Len(t, cfg.configs, 2)
	assert.Contains(t, cfg.configs, "DRVBot")
	assert.NotContains(t, cfg.configs, "Hase")
}

func TestCreateNewChatBot(t *testing.T) {
	//Given
	iBot := &BaseBot{
		FirstName: "Walter",
		LastName:  "Riester",
		UserName:  "RWalterBot",
		ID:        490569313,
		Token:     "490569313:AAEp10AaG9omULlCLEmA_Lp8QYhVchGvtgQ",
	}
	//When
	bot, err := CreateNewChatBot()

	//Then
	assert.NoError(t, err)
	assert.Equal(t, iBot, bot)
}
