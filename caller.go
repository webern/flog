// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package flog

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

// normally skip = 2
// Caller uses the callstack to
func Caller(skip int) string {
	pcuintptr, filestring, lineint, okbool := runtime.Caller(skip)

	if okbool {
		fun := runtime.FuncForPC(pcuintptr)
		theName := fun.Name()

		if len(truncateFilepathsByLastIndexof) > 0 {
			fileix := strings.LastIndex(filestring, truncateFilepathsByLastIndexof)
			if fileix > 0 {
				filestring = filestring[fileix:]
			}
		}

		funcix := strings.LastIndex(theName, ".")
		if funcix > 0 {
			theName = theName[funcix+1:]
		}

		b := bytes.Buffer{}
		b.WriteString(filestring)
		b.WriteString(":")
		b.WriteString(strconv.Itoa(lineint))

		if funcix > 0 {
			b.WriteString(" ")
			b.WriteString(theName)
		}

		return b.String()
	}

	return "error could not find the caller"
}
