package chat

import "testing"

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
