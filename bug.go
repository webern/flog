// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import (
	"github.com/sirupsen/logrus"
)

func Bug() {

	doPanic := logrus.GetLevel() == logrus.TraceLevel

	caller := Caller(2)
	str := caller + " - bug"

	if doPanic {
		logrus.Panic(str)
	} else {
		logrus.Error(str)
	}
}

func Bugf(messageFmt string, args ...interface{}) {

	doPanic := logrus.GetLevel() == logrus.TraceLevel

	caller := Caller(2)
	str := caller + " - bug - " + messageFmt

	if doPanic {
		logrus.Panicf(str, args...)
	} else {
		logrus.Errorf(str, args...)
	}
}

func BugMessage(messageFmt string) {

	doPanic := logrus.GetLevel() == logrus.TraceLevel

	caller := Caller(2)
	str := caller + " - bug - " + messageFmt

	if doPanic {
		logrus.Panicf(str)
	} else {
		logrus.Errorf(str)
	}
}
