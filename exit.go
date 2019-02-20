// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Exit(start time.Time) {
	if GetLevel() != TraceLevel {
		return
	}
	caller := Caller(2)
	logrus.Tracef("%s - exited after %s", caller, time.Since(start).String())
}
