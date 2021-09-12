// Copyright (c) 2021 by Matthew James Briggs, https://github.com/webern

package flog

import "testing"

func TestFlog(t *testing.T) {
	SetTruncationPath("flog/")
	Infof("Testing 123: %s", "foo")
}
