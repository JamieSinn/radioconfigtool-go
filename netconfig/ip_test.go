package netconfig

import "testing"

func TestPing(t *testing.T) {
	result := Ping("8.8.8.8")
	if !result {
		t.Fail()
	}
}
