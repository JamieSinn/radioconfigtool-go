package netconfig

import "testing"

func TestPing(t *testing.T) {
	result, err := Ping("8.8.8.8")
	// If the test errors, meaning that it could not create the socket (As it requires admin) then skip it
	if err != nil {
		t.Skip()
	}
	if !result {
		t.Fail()
	}
}
