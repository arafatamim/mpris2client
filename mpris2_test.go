package mpris2client

import (
	"testing"
)

func TestMicroSecond(t *testing.T) {
	v := µsToString(300040000)
	if v != "5:00" {
		t.Error()
	}
}
