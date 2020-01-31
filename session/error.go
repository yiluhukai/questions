package session

import "errors"

var (
	ErrorSessionNotExist    = errors.New("session doesn't exist")
	ErrorSessionKeyNotExist = errors.New("key of session doesn't exist")

)
