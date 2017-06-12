package imaging

import "testing"

func TestIsValid(t *testing.T) {
	/*
	Router reply:
	Line0: Version of router image
	Line1: Hash of version
	Line2: Image build timestamp
	Line3: Image build timestamp Hash
	Line4: Config timestamp
	Line5: Config timestamp hash
	 */
	reply := []string{
		"1.0.0",
		"47cd76e43f74bbc2e1baaf194d07e1fa",
		"1497291398",
		"b4712c28fc51ff04d8464888be87bbd9",
		"1497291399",
		"759462ea32fd38b55c1ba955f5df819b",
	}
	str := ""
	for _, s := range reply {
		str += s + "\n"
	}
	res := IsValid(str)
	if !res {
		t.Fail()
	}
}


func TestIsInValid(t *testing.T) {
	/*
	Router reply:
	Line0: Version of router image
	Line1: Hash of version
	Line2: Image build timestamp
	Line3: Image build timestamp Hash
	Line4: Config timestamp
	Line5: Config timestamp hash
	 */
	reply := []string{
		"1.0.0",
		"47cd76e43f74bbc2e1baaf194d07e1fa",
		"1497291398",
		"b4712c28fc51ff04d8464888be87bbd9",
		"1497291399",
		// Invalid hash
		"759462ea32fd38b55c1ba955f5df8192",
	}
	str := ""
	for _, s := range reply {
		str += s + "\n"
	}
	res := IsValid(str)
	if res {
		t.Fail()
	}
}
