package errors

import "errors"

func Wrap(err error, message string) error {
	return errors.New(message + ": " + err.Error())
}
