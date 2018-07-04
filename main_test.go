package main

import (
	"testing"
)

func TestReverseMessage(t *testing.T) {
	// given
	var cases = []struct {
		msg        string // input
		reverseMsg string // expected result
	}{
		{"Hallo", "ollaH"},
		{"Hello World!", "!dlroW olleH"},
		{"Fun ", "nuF"},               // trailing whitespaces are removed
		{" No way!", "!yaw oN"},       // leading whitespaces are removed
		{"    ", ""},                  // one or more whitespaces without any char a treated as no message at all
		{"  Anna Abba ", "abbA annA"}, // only leading and trailing whitespaces are removed
	}

	for i, c := range cases {
		// when
		got := ReverseMessage(c.msg)
		// then
		if got != c.reverseMsg {
			t.Errorf("Case %v : got %s but want %s", i, got, c.reverseMsg)
		}
	}
}

func TestProcessingUserRequest(t *testing.T) {
	// given
	var cases = []struct {
		request map[string]string // input
		reply   string            // expected result
	}{
		{map[string]string{"rev": ""}, "Hello World!"},                        //known cmd but no arg
		{map[string]string{"": ""}, "Hello World!"},                           //unknown cmd and no arg
		{map[string]string{"rev": "hallo"}, "ollah"},                          //reverse arg
		{map[string]string{"": "hallo"}, "hallo"},                             //echoing message
		{map[string]string{"revv": "hallo"}, "Sorry, unknown command: /revv"}, //unknown cmd
	}

	for _, c := range cases {
		// when
		got := ProcessingUserRequest(c.request)
		// then
		if got != c.reply {
			t.Errorf("Want %v but got %v", c.reply, got)
		}
	}
}

// TODO: Why time out? Comment in and analyse stack trace!
// src : https://www.hugopicado.com/2016/10/01/testing-over-golang-channels.html
// func TestChatHandler(t *testing.T) {
// 	// given
// 	// for signaling
// 	done := make(chan bool)
// 	defer close(done)

// 	userRq := make(chan chat.Message)
// 	defer close(userRq)

// 	userFb := make(chan chat.Message)

// 	chatMsg := chat.Message{
// 		ChatID:   1234,
// 		UserID:   1234,
// 		UserName: "test",
// 		Request:  map[string]string{"rev": "hallo"},
// 		Reply:    "",
// 	}

// 	// input write routine
// 	go func() {
// 		userRq <- chatMsg
// 		// signaling that input writing is done
// 		done <- true
// 	}()

// 	// when
// 	ChatHandler(userRq, userFb)
// 	<-done // blocks until the input write routine is finished

// 	// then
// 	expected := ReverseMessage(chatMsg.Request["rev"])
// 	found := <-userFb // blocks until the output has contents

// 	if found.Reply != expected {
// 		t.Errorf("Want %v but got %v", expected, found.Reply)
// 	}

// }
