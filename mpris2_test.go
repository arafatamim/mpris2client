package mpris2client

import (
	"testing"
)

func TestMicroSecond(t *testing.T) {
	fiveMins := µsToString(3e+8)
	thirtyMins := µsToString(1.8e+9)
	oneHour := µsToString(3.6e+9)
	// oneDay := µsToString(8.64e+10)
	// println(fiveMins)
	// println(thirtyMins)
	// println(oneHour)
	// println(oneDay)
	if fiveMins != "5:00" {
		t.Error()
	}
	if thirtyMins != "30:00" {
		t.Error()
	}
	if oneHour != "1:00:00" {
		t.Error()
	}
}
