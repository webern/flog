// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import "github.com/sirupsen/logrus"

// NotImplemented can be used as a placeholder in a method stub
func NotImplemented() {
	if logrus.GetLevel() >= logrus.FatalLevel {
		caller := Caller(2)
		str := caller + " - not implemented"
		logrus.Fatal(str)
	}
}

// Todo can be used in a method so that you get reminders in the log that you have unfinished work
func Todo(message string) {
	if logrus.GetLevel() >= logrus.WarnLevel {
		caller := Caller(2)
		str := caller + " - TODO - " + message
		logrus.Warn(str)
	}
}
