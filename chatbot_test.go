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
}
