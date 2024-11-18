package errutil

import "errors"

var ErrRecordNotFound = errors.New("record not found")

type Error struct {
	Message       string
	Type          string
	originalError error
}

func (e *Error) Error() string {
	return e.Message
}

func NewInputError(err error) *Error {
	return &Error{
		originalError: err,
		Type:          "INPUT",
	}
}
