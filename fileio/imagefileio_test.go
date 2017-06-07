package fileio

import (
	"testing"
	"firstinspires.org/radioconfigtool/resources"
)

func TestVerifyImageOM5PAC(t *testing.T) {
	data, err := resources.Asset("firmwareOM5PAC.bin")

	if err != nil {
		t.Fail()
	}

	if !VerifyImage(data, RouterImage{}, 3180218) {
		t.Fail()
	}
}

func TestVerifyImageOM5PAN(t *testing.T) {
	data, err := resources.Asset("firmwareOM5PAN.bin")

	if err != nil {
		t.Fail()
	}

	if !VerifyImage(data, RouterImage{}, 3180361) {
		t.Fail()
	}
}
