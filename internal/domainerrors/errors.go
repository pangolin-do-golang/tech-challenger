package domainerrors

import "errors"

var ErrRecordNotFound = errors.New("record not found")

type Error struct {
	message       string
	originalError error
	isInternal    bool
}

func (e *Error) Error() string {
	return e.message
}
