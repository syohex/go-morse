package morse

import (
	"testing"
)

func TestEncode(t *testing.T) {
	got := Encode("AAA", EN)
	if got != ".- .- .-" {
		t.Errorf("got='%s', expected='.- .- .-'\n", got)
	}
}

func TestDecode(t *testing.T) {
	got := Decode(".- .- .-", EN)
	if got != "AAA" {
		t.Errorf("got='%s', expected='AAA'\n", got)
	}
}
