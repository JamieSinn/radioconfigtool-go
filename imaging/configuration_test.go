package imaging

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	/*
	Router reply:
	Line0: Model: {Hardware Enum}
	Line1: Version: {OpenWRT Build Version}
	Line2: Event: {EventInfo}
	 */
	reply := []string{
		"Model: 1",
		"Version: 2017.2",
		"Event: ",
	}

	res := IsValid(reply)
	if !res {
		t.Fail()
	}
}

func TestIsInValid(t *testing.T) {
	/*
	Router reply:
	Line0: Model: {Hardware Enum}
	Line1: Version: {OpenWRT Build Version}
	Line2: Event: {EventInfo}
	 */
	reply := []string{
		"Model: 1",
		"Version: 2017.2",
		// No Space
		"Event:",
	}

	res := IsValid(reply)
	if res {
		t.Fail()
	}
}

func TestIsUpToDate(t *testing.T) {
	/*
	Router reply:
	Line0: Model: {Hardware Enum}
	Line1: Version: {OpenWRT Build Version}
	Line2: Event: {EventInfo}
	 */
	reply := []string{
		"Model: 1",
		"Version: 2017.2",
		"Event: ",
	}

	res := IsUpToDate(reply[1])
	if !res {
		t.Fail()
	}
}

func TestIsOutOfDate(t *testing.T) {
	/*
	Router reply:
	Line0: Model: {Hardware Enum}
	Line1: Version: {OpenWRT Build Version}
	Line2: Event: {EventInfo}
	 */
	reply := []string{
		"Model: 1",
		"Version: 2017.1",
		"Event: ",
	}

	res := IsUpToDate(reply[1])
	if res {
		t.Fail()
	}
}
