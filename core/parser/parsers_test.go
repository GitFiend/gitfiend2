package parser

import (
	"gitfiend2/core/input"
	"testing"
)

func TestChar(t *testing.T) {
	cParser := Char('c')

	res := Parse(cParser, "c")

	if res.Value != 'c' {
		t.Errorf("Expected c, got %c", res.Value)
	}

	res2 := Parse(cParser, "d")

	if res2.Value != rune(0) {
		t.Error("Expected empty result, got ", res2.Value)
	}

	t.Run(
		`Fails when there's no input (rather than panic)'`, func(t *testing.T) {
			res := Parse(cParser, "")

			if !res.Failed {
				t.Error(`Expected parse failure with no input`)
			}
		},
	)
}

func TestMapOnCharParser(t *testing.T) {
	dParser := Map(
		Char('d'), func(result rune) string {
			return string(result)
		},
	)

	res := Parse(dParser, "d")

	if res.Value != "d" {
		t.Error("Expected \"d\" , got ", res)
	}
}

func TestWord(t *testing.T) {
	t.Run(
		`Parse omg`, func(t *testing.T) {
			wParser := Word("omg")

			res := Parse(wParser, "omg")

			if res.Value != "omg" {
				t.Error("Expected \"omg\", got \"", res, "\"")
			}
		},
	)

	t.Run(
		"Word parsing doesn't go out of bounds", func(t *testing.T) {
			res := Parse(Word("omgg"), "omg")

			if !res.Failed {
				t.Error("Expected this to fail (also not panic).")
			}
		},
	)
}

func TestOptionalWhiteSpace(t *testing.T) {
	t.Run(
		"Parses single space", func(t *testing.T) {
			res := Parse(Ws, " ")

			if res.Failed {
				t.Error("Whitespace parser should always succeed")
			}
		},
	)

	t.Run(
		"Parses single space and then another parser", func(t *testing.T) {
			res := Parse(And2(Ws, Char('c')), " c")

			if res.Value.A != "" {
				t.Error("Failed to get whitespace result")
			}

			if res.Value.B != 'c' {
				t.Error("Failed to get c")
			}
		},
	)

	t.Run(
		"Gets to the end of input and doesn't panic", func(t *testing.T) {
			res := Parse(Ws, "   ")

			if res.Failed {
				t.Error(`White space parser failed on "   "`)
			}
		},
	)
}

func TestRepParserSep(t *testing.T) {
	res := Parse(RepParserSep(Char('a'), Char(',')), "a,a,a")

	if res.Failed {
		t.Error(`Failed to parse "a,a,a"`)
	}
}

func TestUntil(t *testing.T) {
	t.Run(
		`Returns everything before "omg"`, func(t *testing.T) {
			in := input.Input{Code: "abcdefghijklmnomg"}
			res := Until("omg")(&in)

			if res.Failed {
				t.Error("Expected success")
			}
			if res.Value != "abcdefghijklmn" {
				t.Error(`Expected "abcdefghijklmn"`)
			}
			if !in.End() {
				t.Error(`Expected end of input`)
			}
		},
	)

	t.Run(
		`Doesn't go out of bounds if not found`, func(t *testing.T) {
			in := input.Input{Code: "abcdefghijklmn"}
			res := Until("omg")(&in)

			if !res.Failed {
				t.Error(`Expected failure as string isn't in input'`)
			}
			if in.Position != 0 {
				t.Error(`Position should go back to start (0) when parsing fails, got `, in.Position)
			}
		},
	)
}

func TestMany(t *testing.T) {
	res := Parse(Many(Char('c')), "ccc")

	if res.Failed {
		t.Error("Expected success")
	}
	if len(res.Value) != 3 {
		t.Error(`Expected 3 results`)
	}
}
