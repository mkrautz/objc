package foundation

import "testing"

func TestString(t *testing.T) {
	hi := "hello, world!"
	str := NSStringFromString(hi)
	if str.String() != hi {
		t.Error("mismatch")
	}
}
