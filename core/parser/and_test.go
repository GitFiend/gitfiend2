package parser

import "testing"

func TestAnd(t *testing.T) {
	t.Run("When given a Char and Word parser, doesn't fail", func(t *testing.T) {
		p := And2(Char('3'), Word("omg"))

		res, ok := Parse(p, "3omg")

		if !ok {
			t.Error("And2 parser failed.")
		}
		if res.R1 != '3' {
			t.Error("Failed to get first char, got ", res.R1)
		}
		if res.R2 != "omg" {
			t.Error("Failed to get following word omg, got ", res.R2)
		}
	})

	t.Run("When given 3 parsers", func(t *testing.T) {
		p := And3(Char('3'), Word("omg"), Char('o'))

		res, ok := Parse(p, "3omgo")

		if !ok {
			t.Error("And2 parser failed.")
		}
		if res.R1 != '3' {
			t.Error("Failed to get first char, got ", res.R1)
		}
		if res.R2 != "omg" {
			t.Error("Failed to get following word omg, got ", res.R2)
		}
	})

	t.Run("And1", func(t *testing.T) {
		p := And1(Char('1'))

		_, ok := Parse(p, "1")

		if !ok {
			t.Error("And1 Failed")
		}
	})

	t.Run("And4", func(t *testing.T) {
		p := And4(Char('1'), Char('2'), Char('3'), Char('4'))

		_, ok := Parse(p, "1234")

		if !ok {
			t.Error("And4 Failed")
		}
	})

	t.Run("And5", func(t *testing.T) {
		p := And5(Char('1'), Char('2'), Char('3'), Char('4'), Char('5'))

		_, ok := Parse(p, "12345")

		if !ok {
			t.Error("And5 Failed")
		}
	})

	t.Run("And6", func(t *testing.T) {
		_, ok := Parse(And6(Char('1'), Char('2'), Char('3'), Char('4'), Char('5'), Char('6')), "123456")

		if !ok {
			t.Error("And6 Failed")
		}
	})

	t.Run("And7", func(t *testing.T) {
		_, ok := Parse(And7(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'),
		), "1234567")

		if !ok {
			t.Error("And7 Failed")
		}
	})

	t.Run("And8", func(t *testing.T) {
		_, ok := Parse(And8(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
		), "12345678")

		if !ok {
			t.Error("And8 Failed")
		}
	})

	t.Run("And15", func(t *testing.T) {
		_, ok := Parse(And15(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
			Word("11"), Word("12"),
			Word("13"), Word("14"),
			Word("15"),
		), "123456789101112131415")

		if !ok {
			t.Error("And15 Failed")
		}
	})
}
