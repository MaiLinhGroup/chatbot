package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadChatBotCfg(t *testing.T) {
	//Given
	cfg := &ChatBotCfg{}
	iBot := BaseBot{
		FirstName: "Walter",
		LastName:  "Riester",
		UserName:  "RWalterBot",
		ID:        490569313,
		Token:     "490569313:AAEp10AaG9omULlCLEmA_Lp8QYhVchGvtgQ",
	}

	//When
	cfg.ReadChatBotCfg()

	//Then
	assert.Len(t, cfg.configs, 2)
	assert.Contains(t, cfg.configs, "DRVBot")
	assert.NotContains(t, cfg.configs, "Hase")
	assert.Equal(t, iBot, cfg.configs["DRVBot"])
}

func TestGetNewChatBot(t *testing.T) {
	//Given
	iBot := BaseBot{
		FirstName: "Walter",
		LastName:  "Riester",
		UserName:  "RWalterBot",
		ID:        490569313,
		Token:     "490569313:AAEp10AaG9omULlCLEmA_Lp8QYhVchGvtgQ",
	}

	//When
	bot, err := GetNewChatBot(iBot)

	//Then
	assert.NoError(t, err)
	assert.Equal(t, iBot.ID, bot.Self.ID)
}
