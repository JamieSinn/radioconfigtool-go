package imaging

import (
	"testing"
	"firstinspires.org/radioconfigtool/util"
)

func TestEncryptDecrypt(t *testing.T) {
	conf := "7364260261692962810793672227443888951894855411148636427833416688690047096986774713905710607029474025"
	enc := EncryptConfigString(conf)
	util.Debug(enc)
	dec := DecryptConfigString(enc)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}

func TestDecrypt(t *testing.T) {
	conf := "7364260261692962810793672227443888951894855411148636427833416688690047096986774713905710607029474025"
	enc := "Dn7HGlrm8LIwl6+QkY0pzqF4AwZg7g646/xDLXYBBnFk1sDWRrwPnQwOtCn2UWuYDTFAxsoq5RjBfVYhhNuascFpH723zoZzRVDsbtsON863TvRdr9wFH2QYx2ErjdLYuqEm2d6GI5WrLWqNTCvl+3BVbo6MoTiyzyIvGZ3CqAf9AxUfL7JBy3BFIqcgowRcVvC4gnID87E="
	util.Debug(enc)
	dec := DecryptConfigString(enc)
	util.Debug(dec)

	if dec != conf {
		t.Fail()
	}
}

func TestConfigEncryptDecrypt(t *testing.T) {
	conf := "AP5,1334,1334,8charlen,N,N,Y,0,0,,\n"
	util.Debug(conf)
	enc := EncryptConfigString(conf)
	util.Debug(enc)
	dec := DecryptConfigString(enc)
	util.Debug(dec)
	if dec != conf {
		t.Fail()
	}
}
