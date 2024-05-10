package parser

import "testing"

func TestNumberParser(t *testing.T) {
	res := Parse(Number, "123 fd3s")

	if res.Value != "123" {
		t.Error("Expected 123, got ", res)
	}

	res2 := Parse(Number, "-1.009")

	if res2.Value != "-1.009" {
		t.Error("Expected -1.009, got ", res)
	}
}

//func TestUint(t *testing.T) {
//	res := Parse(Uint())
//}
