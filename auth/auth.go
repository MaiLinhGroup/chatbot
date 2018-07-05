package auth

import (
	"errors"
	"os"
	"strconv"
)

// monkey patching :monkey_face:
var osGetEnv = os.Getenv

// User ...
type User struct {
	UserName string
}

// Admin ...
type Admin struct {
	User // an Admin is also a User (embedded Types)
	Key  int
}

// CreateAdminUser ...
func CreateAdminUser() (*Admin, error) {
	key := osGetEnv("ADMIN")
	if key == "" {
		return nil, errors.New("admin key is missing")
	}

	a := new(Admin)

	k, err := strconv.Atoi(key)
	if err != nil {
		return nil, err
	}

	a.Key = k

	return a, nil
}

// Admin ...
func (a *Admin) Admin(key int) bool {
	if key == a.Key {
		return true
	}
	return false
}
