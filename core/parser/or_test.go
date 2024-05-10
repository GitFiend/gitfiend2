package parser

import "testing"

func TestOr(t *testing.T) {
	t.Run("When given one word parser, expect hi result", func(t *testing.T) {
		res := Parse(Or(Word("hi")), "hi")

		if res.Failed {
			t.Error("Expected single word parser to succeed")
		}
	})

	t.Run("When given 2 Char parsers, expect seconde d result", func(t *testing.T) {
		res := Parse(Or(Char('c'), Char('d')), "d")

		if res.Failed {
			t.Error("Failed to parse d")
		}
		if res.Value != 'd' {
			t.Error("Expected second parser to succeed")
		}
	})
}
