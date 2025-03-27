package storages

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExist    = errors.New("user exist")
)
