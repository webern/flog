// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

package flog

import (
	"github.com/sirupsen/logrus"
)

func Bug() {

	doPanic := logrus.GetLevel() == logrus.TraceLevel

	caller := BlazeCaller(2)
	str := caller + " - bug"

	if doPanic {
		logrus.Panic(str)
	} else {
		logrus.Error(str)
	}
}

func Bugf(messageFmt string, args ...interface{}) {

	doPanic := logrus.GetLevel() == logrus.TraceLevel

	caller := BlazeCaller(2)
	str := caller + " - bug - " + messageFmt

	if doPanic {
		logrus.Panicf(str, args...)
	} else {
		logrus.Errorf(str, args...)
	}
}

func BugMessage(messageFmt string) {

	doPanic := logrus.GetLevel() == logrus.TraceLevel

	caller := BlazeCaller(2)
	str := caller + " - bug - " + messageFmt

	if doPanic {
		logrus.Panicf(str)
	} else {
		logrus.Errorf(str)
	}
}
