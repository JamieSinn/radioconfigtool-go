package util

import "testing"

func TestIsValidTeamNumber(t *testing.T) {
	valid := []string{"0001", "1334", "1", "7777", "254", "11", "9999"}
	invalid := []string{"-0001", "11334", "0", "10000", "x", "string"}

	for _, s := range valid {
		result := IsValidTeamNumber(s)
		if !result {
			t.Fail()
		}
	}

	for _, s := range invalid {
		result := IsValidTeamNumber(s)
		if result {
			t.Fail()
		}
	}
}
