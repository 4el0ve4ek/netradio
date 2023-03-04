package errors

import "errors"

func New(text string) error {
	return errors.New(text)
}

func Wrap(err error, message string) error {
	return errors.New(message + ": " + err.Error())
}
