package errors

import (
	"fmt"
)

type causer interface {
	Cause() error
}

type Error struct {
	Code    Code
	Message string
	cause   error
}

func (err Error) Error() string {
	return err.Message
}

func (err Error) Cause() error {
	return err.cause
}

func New(code Code, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

func Errorf(code Code, format string, args ...interface{}) Error {
	return New(code, fmt.Sprintf(format, args...))
}

func Wrap(err error, message string) Error {
	return Error{
		Message: message,
		cause:   err,
	}
}

func Wrapf(err error, format string, args ...interface{}) Error {
	return Wrap(err, fmt.Sprintf(format, args...))
}
