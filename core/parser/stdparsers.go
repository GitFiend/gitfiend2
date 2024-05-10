package parser

import (
	"regexp"
	"unicode"
)

var Number = Regex(regexp.MustCompile(`-?(\d+(\.\d+)?)`))
var Int = Regex(regexp.MustCompile(`^[-+]?\d+`))

var AnyWord = Regex(regexp.MustCompile(`\w+`))

var Uint = uintParser()

func uintParser() Parser[string] {
	return func(in *Input) Result[string] {
		var parts []rune

		for !in.End() {
			n := in.NextChar()

			if unicode.IsDigit(n) {
				parts = append(parts, n)

				in.Advance()
			} else {
				break
			}
		}

		if len(parts) > 0 {
			return Result[string]{Value: string(parts)}
		}

		return Result[string]{Failed: true}
	}
}
