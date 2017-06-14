package imaging

import (
	"testing"
	"strconv"
	"time"
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

func TestIsWithinCompetition(t *testing.T) {
	/*
		Router reply:
		Line0: Model: {Hardware Enum}
		Line1: Version: {OpenWRT Build Version}
		Line2: Event: {EventInfo}
		 */
	reply := []string{
		"Model: 1",
		"Version: 2017.1",
		"Event: " + strconv.FormatInt(time.Now().Unix()+3000, 10),
	}

	res := IsWithinCompetition(reply[2])
	// This test should return true. (It is within competition)
	if !res {
		t.Fail()
	}
}

func TestIsOutsideCompetition(t *testing.T) {
	/*
		Router reply:
		Line0: Model: {Hardware Enum}
		Line1: Version: {OpenWRT Build Version}
		Line2: Event: {EventInfo}
		 */
	reply := []string{
		"Model: 1",
		"Version: 2017.1",
		"Event: "+ strconv.FormatInt(time.Now().Unix()-30000, 10),
	}
	// This test should return false. (It is outside competition)
	res := IsWithinCompetition(reply[2])
	if res {
		t.Fail()
	}
}



func TestIsEmptyEvent(t *testing.T) {
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

	res := IsWithinCompetition(reply[2])
	// Because the team config string has no event, it should return false.
	if res {
		t.Fail()
	}
}


