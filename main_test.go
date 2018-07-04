package main

import "testing"

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
