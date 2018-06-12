package chat

import "testing"

func TestReverseMessage(t *testing.T) {
	// given
	var cases = []struct {
		msg         string // input
		reversedMsg string // expected result
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
		got := ReversedMessage(c.msg)
		// then
		if got != c.reversedMsg {
			t.Errorf("Case %v : got %s but want %s", i, got, c.reversedMsg)
		}
	}
}

// TestNewNoToken tests the case that the token is not set in the environment so that a new error is expected
func TestNewNoToken(t *testing.T) {
	// given
	tmpOSGetEnv := osGetEnv
	defer func() {
		osGetEnv = tmpOSGetEnv
	}()
	// providing a new behaviour for Getenv call
	osGetEnv = func(key string) string {
		return ""
	}
	// when
	_, err := New()
	// then
	if err == nil {
		t.Errorf("Want error because token was not set, but got %v", err)
	}
}
