package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleUserRequest(t *testing.T) {
	//given
	updateRetriever = func() Update { return Update{from: "MLEdith", text: "test 1", chatID: 480821480} }

	exp := UserRequest{
		chatID: 480821480,
		msg:    "test 1",
		cmds:   make(map[string]string),
	}

	exp.cmds["help"] = ""

	//when
	subj := HandleUserRequest()

	//then
	assert.Equal(t, exp.chatID, subj.chatID)
	assert.Equal(t, exp.msg, subj.msg)
}
