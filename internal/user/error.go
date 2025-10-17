package user

import "errors"

var ErrFirstNameRequiered = errors.New("first name is requiered")
var ErrLastNameRequiered = errors.New("last name is requiered")

var ErrUserNotfound = errors.New("user not found")
