// iqBid Pricing System, Copyright (c) 2019 by Inteliquent, Inc.

package flog

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

// normally skip = 2
func BlazeCaller(skip int) string {
	pcuintptr, filestring, lineint, okbool := runtime.Caller(skip)

	if okbool {
		fun := runtime.FuncForPC(pcuintptr)
		theName := fun.Name()

		fileix := strings.LastIndex(filestring, "blaze/")
		if fileix > 0 {
			filestring = filestring[fileix:]
		}

		funcix := strings.LastIndex(theName, ".")
		if funcix > 0 {
			theName = theName[funcix+1:]
		}

		b := bytes.Buffer{}
		b.WriteString(filestring)
		b.WriteString(" (")
		b.WriteString(strconv.Itoa(lineint))
		b.WriteString(")")

		if funcix > 0 {
			b.WriteString(": ")
			b.WriteString(theName)
		}

		return b.String()
	}

	return "error could not find the caller"
}
