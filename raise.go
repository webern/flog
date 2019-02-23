// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import (
	"errors"
	"fmt"
	"strconv"
)

func Raise(message string) error {
	caller := Caller(2)
	str := caller + " - error - " + message
	return errors.New(str)
}

func Raisef(messageFmt string, args ...interface{}) error {
	caller := Caller(2)
	str := caller + " - error - " + messageFmt
	return fmt.Errorf(str, args...)
}

func RaiseBadJobID(jobID int) error {
	caller := Caller(2)
	str := caller + " - bad job id - " + strconv.Itoa(jobID)
	return errors.New(str)
}
