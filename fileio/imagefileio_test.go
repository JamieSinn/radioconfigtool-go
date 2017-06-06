package fileio

import (
	"testing"
	"io/ioutil"
)

func TestVerifyImageOM5PAC(t *testing.T) {
	data, err := ioutil.ReadFile("firmwareOM5PAC.bin")

	if err != nil {
		t.Fail()
	}

	if !VerifyImage(data, RouterImage{}, 3180218) {
		t.Fail()
	}
}

func TestVerifyImageOM5PAN(t *testing.T) {
	data, err := ioutil.ReadFile("firmwareOM5PAN.bin")

	if err != nil {
		t.Fail()
	}

	if !VerifyImage(data, RouterImage{}, 3180361) {
		t.Fail()
	}
}
