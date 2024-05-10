package parser

import "testing"

func TestAnd(t *testing.T) {
	t.Run("When given a Char and Word parser, doesn't fail", func(t *testing.T) {
		p := And2(Char('3'), Word("omg"))

		res := Parse(p, "3omg")

		if res.Failed {
			t.Error("And2 parser failed.")
		}
		if res.Value.A != '3' {
			t.Error("Failed to get first char, got ", res.Value.A)
		}
		if res.Value.B != "omg" {
			t.Error("Failed to get following word omg, got ", res.Value.B)
		}
	})

	t.Run("When given 3 parsers", func(t *testing.T) {
		p := And3(Char('3'), Word("omg"), Char('o'))

		res := Parse(p, "3omgo")

		if res.Failed {
			t.Error("And2 parser failed.")
		}
		if res.Value.A != '3' {
			t.Error("Failed to get first char, got ", res.Value.A)
		}
		if res.Value.B != "omg" {
			t.Error("Failed to get following word omg, got ", res.Value.B)
		}
	})

	t.Run("And1", func(t *testing.T) {
		p := And1(Char('1'))

		res := Parse(p, "1")

		if res.Failed {
			t.Error("And1 Failed")
		}
	})

	t.Run("And4", func(t *testing.T) {
		p := And4(Char('1'), Char('2'), Char('3'), Char('4'))

		res := Parse(p, "1234")

		if res.Failed {
			t.Error("And4 Failed")
		}
	})

	t.Run("And5", func(t *testing.T) {
		p := And5(Char('1'), Char('2'), Char('3'), Char('4'), Char('5'))

		res := Parse(p, "12345")

		if res.Failed {
			t.Error("And5 Failed")
		}
	})

	t.Run("And6", func(t *testing.T) {
		res := Parse(And6(Char('1'), Char('2'), Char('3'), Char('4'), Char('5'), Char('6')), "123456")

		if res.Failed {
			t.Error("And6 Failed")
		}
	})

	t.Run("And7", func(t *testing.T) {
		res := Parse(And7(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'),
		), "1234567")

		if res.Failed {
			t.Error("And7 Failed")
		}
	})

	t.Run("And8", func(t *testing.T) {
		res := Parse(And8(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
		), "12345678")

		if res.Failed {
			t.Error("And8 Failed")
		}
	})

	t.Run("And9", func(t *testing.T) {
		res := Parse(And9(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'),
		), "123456789")

		if res.Failed {
			t.Error("And9 Failed")
		}
	})

	t.Run("And10", func(t *testing.T) {
		res := Parse(And10(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
		), "12345678910")

		if res.Failed {
			t.Error("And10 Failed")
		}
	})

	t.Run("And11", func(t *testing.T) {
		res := Parse(And11(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
			Word("11"),
		), "1234567891011")

		if res.Failed {
			t.Error("And11 Failed")
		}
	})

	t.Run("And12", func(t *testing.T) {
		res := Parse(And12(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
			Word("11"), Word("12"),
		), "123456789101112")

		if res.Failed {
			t.Error("And12 Failed")
		}
	})

	t.Run("And13", func(t *testing.T) {
		res := Parse(And13(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
			Word("11"), Word("12"),
			Word("13"),
		), "12345678910111213")

		if res.Failed {
			t.Error("And13 Failed")
		}
	})

	t.Run("And14", func(t *testing.T) {
		res := Parse(And14(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
			Word("11"), Word("12"),
			Word("13"), Word("14"),
		), "1234567891011121314")

		if res.Failed {
			t.Error("And14 Failed")
		}
	})

	t.Run("And15", func(t *testing.T) {
		res := Parse(And15(
			Char('1'), Char('2'),
			Char('3'), Char('4'),
			Char('5'), Char('6'),
			Char('7'), Char('8'),
			Char('9'), Word("10"),
			Word("11"), Word("12"),
			Word("13"), Word("14"),
			Word("15"),
		), "123456789101112131415")

		if res.Failed {
			t.Error("And15 Failed")
		}
	})
}
