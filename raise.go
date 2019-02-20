// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

package flog

import (
	"errors"
	"fmt"
	"strconv"
)

func Raise(message string) error {
	caller := BlazeCaller(2)
	str := caller + " - error - " + message
	return errors.New(str)
}

func Raisef(messageFmt string, args ...interface{}) error {
	caller := BlazeCaller(2)
	str := caller + " - error - " + messageFmt
	return fmt.Errorf(str, args...)
}

func RaiseBadJobID(jobID int) error {
	caller := BlazeCaller(2)
	str := caller + " - bad job id - " + strconv.Itoa(jobID)
	return errors.New(str)
}
