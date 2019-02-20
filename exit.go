// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

package flog

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Exit(start time.Time) {
	if GetLevel() != logrus.TraceLevel {
		return
	}
	caller := BlazeCaller(2)
	logrus.Tracef("%s - exited after %s", caller, time.Since(start).String())
}
