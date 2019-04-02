// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import (
	"github.com/sirupsen/logrus"
)

// LogErr can wrap a function that returns an error in cases where you do not want to handle ther error.
// for example 'defer LogErr(file.Close())'
func LogErr(err error) {
	if err == nil {
		return
	}

	caller := Caller(2)
	logrus.Errorf("%s - unhandled error: %s", caller, err.Error())
}

// LogErr2 can wrap a function that returns an error and something else we don't care about
func LogErr2(ignored interface{}, err error) {
	if err == nil {
		return
	}

	caller := Caller(2)
	logrus.Errorf("%s - unhandled error: %s", caller, err.Error())
}

// LogErr2 can wrap a function that returns an error and something else we don't care about
func LogErr3(ignored1, ignored2 interface{}, err error) {
	if err == nil {
		return
	}

	caller := Caller(2)
	logrus.Errorf("%s - unhandled error: %s", caller, err.Error())
}
