package main

import "testing"

func TestNewChatBot(t *testing.T) {
	cb := &ChatBot{
		bot: &telegrambotapi{},
	}
	cb.NewChatBot()
	if cb.bot == nil {
		t.Error("Expected *telegrambotapi, but got nil")
	}
}
