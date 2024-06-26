package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

func TestChar(t *testing.T) {
	cParser := Rune('c')

	res, _ := ParseAll(cParser, "c")

	if res != 'c' {
		t.Errorf("Expected c, got %c", res)
	}

	res2, _ := ParseAll(cParser, "d")

	if res2 != rune(0) {
		t.Error("Expected empty result, got ", res2)
	}

	t.Run(
		`Fails when there's no input (rather than panic)'`, func(t *testing.T) {
			_, ok := ParseAll(cParser, "")

			if ok {
				t.Error(`Expected parse failure with no input`)
			}
		},
	)
}

func TestMapOnCharParser(t *testing.T) {
	dParser := Map(
		Rune('d'), func(result rune) string {
			return string(result)
		},
	)

	res, _ := ParseAll(dParser, "d")

	if res != "d" {
		t.Error("Expected \"d\" , got ", res)
	}
}

func TestWord(t *testing.T) {
	t.Run(
		`ParseAll omg`, func(t *testing.T) {
			wParser := Word("omg")

			res, _ := ParseAll(wParser, "omg")

			if res != "omg" {
				t.Error("Expected \"omg\", got \"", res, "\"")
			}
		},
	)

	t.Run(
		"Word parsing doesn't go out of bounds", func(t *testing.T) {
			_, ok := ParseAll(Word("omgg"), "omg")

			if ok {
				t.Error("Expected this to fail (also not panic).")
			}
		},
	)
}

func TestOptionalWhiteSpace(t *testing.T) {
	t.Run(
		"Parses single space", func(t *testing.T) {
			_, ok := ParseAll(Ws, " ")

			if !ok {
				t.Error("Whitespace parser should always succeed")
			}
		},
	)

	t.Run(
		"Parses single space and then another parser", func(t *testing.T) {
			res, _ := ParseAll(And2(Ws, Rune('c')), " c")

			if res.R1 != "" {
				t.Error("Failed to get whitespace result")
			}

			if res.R2 != 'c' {
				t.Error("Failed to get c")
			}
		},
	)

	t.Run(
		"Gets to the end of input and doesn't panic", func(t *testing.T) {
			_, ok := ParseAll(Ws, "   ")

			if !ok {
				t.Error(`White space parser failed on "   "`)
			}
		},
	)
}

func TestRepParserSep(t *testing.T) {
	_, ok := ParseAll(RepParserSep(Rune('a'), Rune(',')), "a,a,a")

	if !ok {
		t.Error(`Failed to parse "a,a,a"`)
	}
}

func TestUntil(t *testing.T) {
	t.Run(
		`Returns everything before "omg"`, func(t *testing.T) {
			in := NewInput("abcdefghijklmnomg")
			res, ok := UntilString("omg")(&in)

			if !ok {
				t.Error("Expected success")
			}
			if res != "abcdefghijklmn" {
				t.Error(`Expected "abcdefghijklmn"`)
			}
			if !in.End() {
				t.Error(`Expected end of input`)
			}
		},
	)

	t.Run(
		`Doesn't go out of bounds if not found`, func(t *testing.T) {
			in := NewInput("abcdefghijklmn")
			_, ok := UntilString("omg")(&in)

			if ok {
				t.Error(`Expected failure as string isn't in input'`)
			}
			if in.Position != 0 {
				t.Error(`Position should go back to start (0) when parsing fails, got `, in.Position)
			}
		},
	)
}

func TestTakeRuneWhile(t *testing.T) {
	p := TakeRuneWhile(func(r rune) bool {
		return unicode.IsLetter(r)
	})

	res, ok := ParsePart(p, "abcd55")

	assert.True(t, ok)
	assert.Equal(t, "abcd", res)

	_, ok = ParseAll(p, "abcd55")
	assert.False(t, ok)
}

func TestMany(t *testing.T) {
	res, ok := ParseAll(Many(Rune('c')), "ccc")

	if !ok {
		t.Error("Expected success")
	}
	if len(res) != 3 {
		t.Error(`Expected 3 results`)
	}
}

func TestNotParser(t *testing.T) {
	parser := Not(Word("hello"))

	_, ok := ParsePart(parser, "hi")
	assert.True(t, ok)
}
