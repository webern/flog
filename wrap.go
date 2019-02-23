package flog

import "errors"

func Wrap(err error) error {
	caller := Caller(2)
	str := caller + " - error - " + err.Error()
	return errors.New(str)
}
