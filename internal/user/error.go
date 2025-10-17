package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequiered = errors.New("first name is requiered")
var ErrLastNameRequiered = errors.New("last name is requiered")

// var ErrUserNotfound = errors.New("user not found")

type ErrNotFound struct {
	UserID string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("user '%s' doesn't exist", e.UserID)
}
