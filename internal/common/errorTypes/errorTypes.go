package errorTypes

import "errors"

var (
	ErrUserNotFound = errors.New("user with given id not found")
)
