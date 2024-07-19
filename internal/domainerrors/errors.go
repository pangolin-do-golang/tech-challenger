package domainerrors

import "errors"

var ErrRecordNotFound = errors.New("record not found")

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

type Error struct {
	Message       string
	Type          string
	originalError error
}

func (e *Error) Error() string {
	return e.Message
}

func NewSystemError(err error, mes string) *Error {
	return &Error{
		Message:       mes,
		Type:          "SYSTEM",
		originalError: err,
	}
}

func NewBusinessError(err error, mes string) *Error {
	return &Error{
		Message:       mes,
		originalError: err,
		Type:          "BUSINESS",
	}
}

func NewInputError(err error, mes string) *Error {
	return &Error{
		Message:       mes,
		originalError: err,
		Type:          "INPUT",
	}
}
