package parser

import (
	"unicode"
)

var Ws = OptionalWhiteSpace[string]()

var SignedInt = Map(
	And2(Or(Word("-"), Word("+"), Ws), UInt),
	func(result And2Result[string, string]) string {
		return result.R1 + result.R2
	},
)

var Float = Map(
	And4(Optional(Rune('-')), UInt, Optional(Rune('.')), Optional(UInt)),
	func(res And4Result[rune, string, rune, string]) string {
		var str string
		if res.R1 == '-' {
			str = string(res.R1) + res.R2
		} else {
			str = res.R2
		}

		if res.R3 == '.' {
			str += "." + res.R4
		}
		return str
	},
)

var AnyWord = TakeRuneWhile(
	func(r rune) bool {
		return unicode.IsDigit(r) || unicode.IsLetter(r)
	},
)

var UInt = TakeRuneWhile(
	func(r rune) bool {
		return unicode.IsDigit(r)
	},
)

var LineEnd = Or(Word("\n"), Word("\r\n"))

var UntilLineEnd = UntilParser(LineEnd)

// StringLiteral Note: Naive, doesn't handle escaped quotes.
var StringLiteral = Map(
	And3(
		Rune('"'),
		TakeRuneWhile(
			func(r rune) bool {
				return r != '"'
			},
		),
		Rune('"'),
	), func(res And3Result[rune, string, rune]) string {
		return res.R2
	},
)

var UntilNul = UntilParser(
	ConditionalRune(
		func(r rune) bool {
			return unicode.IsControl(r) && !unicode.IsSpace(r)
		},
	),
)
