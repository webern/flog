// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import "github.com/sirupsen/logrus"

// Must terminates the program if err is not nul
func Must(err error) {
	if err != nil {
		msg := Caller(2) + " - " + err.Error()
		logrus.Fatal(msg)
	}
}

// Must2 terminates the program if err is not nul
func Must2(_ interface{}, err error) {
	if err != nil {
		msg := Caller(2) + " - " + err.Error()
		logrus.Fatal(msg)
	}
}

// Must3 terminates the program if err is not nul
func Must3(_, _ interface{}, err error) {
	if err != nil {
		msg := Caller(2) + " - " + err.Error()
		logrus.Fatal(msg)
	}
}
