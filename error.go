package errors

import (
	"fmt"
	"strings"
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
	if err.cause == nil {
		return err.Message
	}
	causes := []string{}
	func(c error) {
		for {
			if c == nil {
				break
			}
			causes = append(causes, c.Error())
			if err, ok := c.(causer); ok {
				c = err.Cause()
			} else {
				c = nil
			}
		}
	}(err)
	return strings.Join(causes, "\n\t")
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
