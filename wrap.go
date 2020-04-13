package flog

import "errors"

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	caller := Caller(2)
	str := caller + " - error - " + err.Error()
	return errors.New(str)
}
