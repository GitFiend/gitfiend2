package parser

import "testing"

func TestAnd(t *testing.T) {
	t.Run("When given a Rune and Word parser, doesn't fail", func(t *testing.T) {
		p := And2(Rune('3'), Word("omg"))

		res, ok := ParseAll(p, "3omg")

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
		p := And3(Rune('3'), Word("omg"), Rune('o'))

		res, ok := ParseAll(p, "3omgo")

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
		p := And1(Rune('1'))

		_, ok := ParseAll(p, "1")

		if !ok {
			t.Error("And1 Failed")
		}
	})

	t.Run("And4", func(t *testing.T) {
		p := And4(Rune('1'), Rune('2'), Rune('3'), Rune('4'))

		_, ok := ParseAll(p, "1234")

		if !ok {
			t.Error("And4 Failed")
		}
	})

	t.Run("And5", func(t *testing.T) {
		p := And5(Rune('1'), Rune('2'), Rune('3'), Rune('4'), Rune('5'))

		_, ok := ParseAll(p, "12345")

		if !ok {
			t.Error("And5 Failed")
		}
	})

	t.Run("And6", func(t *testing.T) {
		_, ok := ParseAll(And6(Rune('1'), Rune('2'), Rune('3'), Rune('4'), Rune('5'), Rune('6')), "123456")

		if !ok {
			t.Error("And6 Failed")
		}
	})

	t.Run("And7", func(t *testing.T) {
		_, ok := ParseAll(And7(
			Rune('1'), Rune('2'),
			Rune('3'), Rune('4'),
			Rune('5'), Rune('6'),
			Rune('7'),
		), "1234567")

		if !ok {
			t.Error("And7 Failed")
		}
	})

	t.Run("And8", func(t *testing.T) {
		_, ok := ParseAll(And8(
			Rune('1'), Rune('2'),
			Rune('3'), Rune('4'),
			Rune('5'), Rune('6'),
			Rune('7'), Rune('8'),
		), "12345678")

		if !ok {
			t.Error("And8 Failed")
		}
	})

	t.Run("And15", func(t *testing.T) {
		_, ok := ParseAll(And15(
			Rune('1'), Rune('2'),
			Rune('3'), Rune('4'),
			Rune('5'), Rune('6'),
			Rune('7'), Rune('8'),
			Rune('9'), Word("10"),
			Word("11"), Word("12"),
			Word("13"), Word("14"),
			Word("15"),
		), "123456789101112131415")

		if !ok {
			t.Error("And15 Failed")
		}
	})
}
