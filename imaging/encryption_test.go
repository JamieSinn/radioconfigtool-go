package imaging

import (
	"testing"
	"firstinspires.org/radioconfigtool/util"
	"strings"
)

func TestEncryptDecrypt(t *testing.T) {
	conf := "AP5,1334,1334,8charlen,N,N,Y,0,0,,\n"
	enc := EncryptConfigString(conf)
	util.Debug(enc)
	dec := DecryptConfigString(enc)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}


func TestEncryptDecryptCBC_1334(t *testing.T) {
	conf := "AP5,1334,1334,8charlen,N,N,Y,0,0,,\n"
	enc := EncryptConfigStringCBC(conf)
	util.Debug(enc)
	dec := DecryptConfigStringCBC(enc)
	util.Debug(dec)
	dec = strings.Replace(dec, "=", "", -1)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}

func TestEncryptDecryptCBC_1114(t *testing.T) {
	conf := "AP24,1114,1114,8charlen,N,N,Y,0,0,,\n"
	enc := EncryptConfigStringCBC(conf)
	util.Debug(enc)
	dec := DecryptConfigStringCBC(enc)
	util.Debug(dec)
	dec = strings.Replace(dec, "=", "", -1)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}

func TestEncryptDecryptCBC_254(t *testing.T) {
	conf := "B5,254,254,8charlen,N,N,Y,0,0,,\n"
	enc := EncryptConfigStringCBC(conf)
	util.Debug(enc)
	dec := DecryptConfigStringCBC(enc)
	util.Debug(dec)
	dec = strings.Replace(dec, "=", "", -1)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}


func TestEncryptDecryptCBC_11(t *testing.T) {
	conf := "B5,11,11,8charlen,N,N,Y,0,0,,\n"
	enc := EncryptConfigStringCBC(conf)
	util.Debug(enc)
	dec := DecryptConfigStringCBC(enc)
	util.Debug(dec)
	dec = strings.Replace(dec, "=", "", -1)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}


func TestEncryptDecryptCBC_1(t *testing.T) {
	conf := "B5,1,1,8charlen,N,N,Y,0,0,,\n"
	enc := EncryptConfigStringCBC(conf)
	util.Debug(enc)
	dec := DecryptConfigStringCBC(enc)
	util.Debug(dec)
	dec = strings.Replace(dec, "=", "", -1)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}

func TestDecryptCFB(t *testing.T) {
	conf := "7364260261692962810793672227443888951894855411148636427833416688690047096986774713905710607029474025"
	enc := "Dn7HGlrm8LIwl6+QkY0pzqF4AwZg7g646/xDLXYBBnFk1sDWRrwPnQwOtCn2UWuYDTFAxsoq5RjBfVYhhNuascFpH723zoZzRVDsbtsON863" +
		"TvRdr9wFH2QYx2ErjdLYuqEm2d6GI5WrLWqNTCvl+3BVbo6MoTiyzyIvGZ3CqAf9AxUfL7JBy3BFIqcgowRcVvC4gnID87E="
	util.Debug(enc)
	dec := DecryptConfigString(enc)
	util.Debug(dec)

	if dec != conf {
		t.Fail()
	}
}

func TestDecryptCBC(t *testing.T) {
	b64 := "RMmuxFCAtV2S0DVG06Rere92PyNJpe0X0l/h36uKAp1D2u1eEP9HvTZb+HNX51AYAe0OoKsiQPjq2tQxVYk+Tg=="
	util.Debug(b64)
	conf := "AP5,1334,1334,8charlen,N,N,Y,0,0,,\n============="
	dec := DecryptConfigStringCBC(b64)
	util.Debug(dec)

	if dec != conf {
		t.Fail()
	}
}
