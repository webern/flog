// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

package flog

import "github.com/sirupsen/logrus"

func NotImplemented() {
	if logrus.GetLevel() >= logrus.FatalLevel {
		caller := BlazeCaller(2)
		str := caller + " - not implemented"
		logrus.Fatal(str)
	}
}

func Todo(message string) {
	if logrus.GetLevel() >= logrus.WarnLevel {
		caller := BlazeCaller(2)
		str := caller + " - TODO - " + message
		logrus.Warn(str)
	}
}
