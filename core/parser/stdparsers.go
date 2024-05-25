package parser

import (
	"unicode"
)

var SignedInt = Map(And2(Or(Word("-"), Word("+"), Ws), UInt),
	func(result And2Result[string, string]) string {
		return result.R1 + result.R2
	},
)

var AnyWord = TakeRuneWhile(func(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
})

var UInt = TakeRuneWhile(func(r rune) bool {
	return unicode.IsDigit(r)
})

var LineEnd = Or(Word("\n"), Word("\r\n"))

var UntilLineEnd = UntilParser(LineEnd)

// StringLiteral Note: Naive, doesn't handle escaped quotes.
var StringLiteral = Map(And3(
	Rune('"'),
	TakeRuneWhile(func(r rune) bool {
		return r != '"'
	}),
	Rune('"'),
), func(res And3Result[rune, string, rune]) string {
	return res.R2
})
