package parser

import (
	"regexp"
	"unicode"
)

var Number = Regex(regexp.MustCompile(`-?(\d+(\.\d+)?)`))
var Int = Regex(regexp.MustCompile(`^[-+]?\d+`))

// var AnyWord = Regex(regexp.MustCompile(`\w+`))

var AnyWord = TakeRuneWhile(func(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
})

var Uint = uintParser()

var LineEnd = Or(Word("\n"), Word("\r\n"))

var UntilLineEnd = UntilParser(LineEnd)

func uintParser() Parser[string] {
	return func(in *Input) (string, bool) {
		var parts []rune

		for !in.End() {
			n := in.NextRune()

			if unicode.IsDigit(n) {
				parts = append(parts, n)

				in.Advance()
			} else {
				break
			}
		}

		if len(parts) > 0 {
			return string(parts), true
		}

		return "", false
	}
}
