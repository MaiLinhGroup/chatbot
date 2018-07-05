package auth

import "testing"

func TestCreateAdminUserNoKey(t *testing.T) {
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
	_, err := CreateAdminUser()

	// then
	if err == nil {
		t.Error("Want error because no key found but error was nil.")
	}
}

func TestCreateAdminUserAtoiInvalidSyntax(t *testing.T) {
	// given
	tmpOSGetEnv := osGetEnv
	defer func() {
		osGetEnv = tmpOSGetEnv
	}()

	// providing a new behaviour for Getenv call
	osGetEnv = func(key string) string {
		return "?"
	}

	// when
	_, err := CreateAdminUser()

	// then
	if err == nil {
		t.Error("Want error because cannot convert string to int but error was nil.")
	}
}

func TestCreateAdminUserWithKey(t *testing.T) {
	// given
	tmpOSGetEnv := osGetEnv
	defer func() {
		osGetEnv = tmpOSGetEnv
	}()

	// providing a new behaviour for Getenv call
	osGetEnv = func(key string) string {
		return "123"
	}

	want := new(Admin)
	want.Key = 123

	// when
	got, err := CreateAdminUser()

	// then
	if got.Key != want.Key {
		t.Errorf("Want %v but got %v", want.Key, got.Key)
	}
	if err != nil {
		t.Errorf("No error expected but got %v", err.Error())
	}
}
