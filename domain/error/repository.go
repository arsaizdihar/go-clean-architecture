package derror

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)
