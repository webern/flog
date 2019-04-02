// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import (
	"io"

	"github.com/sirupsen/logrus"
)

// LogClose can wrap a function that returns an error in cases where you do not want to handle ther error.
// for example 'defer LogErr(file.Close())'
func LogClose(closer io.Closer) {
	err := closer.Close()

	if err == nil {
		return
	}

	caller := Caller(2)
	logrus.Errorf("%s - unhandled io.Closer error: %s", caller, err.Error())
}
