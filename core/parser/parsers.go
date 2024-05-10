package parser

import (
	"gitfiend2/core/input"
	"regexp"
)

func Char(c rune) Parser[rune] {
	return func(in *input.Input) Result[rune] {
		if !in.End() {
			n := in.NextChar()

			if n == c {
				in.Advance()
				return Result[rune]{Value: n}
			}
		}

		return Result[rune]{Failed: true}
	}
}

func Word(word string) Parser[string] {
	return func(in *input.Input) Result[string] {
		p := in.Position

		for _, c := range word {
			if !in.End() && in.NextChar() == c {
				in.Advance()
			} else {
				in.SetPosition(p)
				return Result[string]{Failed: true}
			}
		}

		return Result[string]{Value: word}
	}
}

// Regex
// var re *regexp.Regexp = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
func Regex(re *regexp.Regexp) Parser[string] {
	return func(i *input.Input) Result[string] {
		code := i.Code[i.Position:]
		match := re.FindStringIndex(code)

		if match[0] == 0 {
			start := i.Position
			i.AdvanceBy(match[1])

			return Result[string]{Value: i.Code[start : start+match[1]]}
		}

		return Result[string]{Failed: true}
	}
}

func OptionalWhiteSpace[T any]() Parser[T] {
	return func(in *input.Input) Result[T] {
		for !in.End() && IsWhiteSpace(in.NextChar()) {
			in.Advance()
		}

		return Result[T]{Value: *new(T)}
	}
}

var Ws = OptionalWhiteSpace[string]()

func IsWhiteSpace(val rune) bool {
	return val == ' ' || val == '\n' || val == '\r' || val == '\t'
}

// RepSep Assumes separator has any amount of whitespace around it.
func RepSep[T any](parser Parser[T], separator string) Parser[[]T] {
	sepParser := And3(Ws, Word(separator), Ws)

	return RepParserSep(parser, sepParser)
}

// RepParserSep Doesn't require a result.
func RepParserSep[T any, U any](parser Parser[T], separator Parser[U]) Parser[[]T] {
	return func(in *input.Input) Result[[]T] {
		var results []T

		for !in.End() {
			result := parser(in)
			if result.Failed {
				//if len(results) == 0 {
				//	//return Result[[]T]{Failed: true}
				//	return Result[[]T]{}
				//}

				break
			}
			results = append(results, result.Value)

			if in.End() {
				return Result[[]T]{Value: results}
			}

			sepRes := separator(in)
			if sepRes.Failed {
				break
			}
		}

		return Result[[]T]{Value: results}
	}
}

// Until
// Input is consumed including str, but str is not included in the result.
func Until(str string) Parser[string] {
	return func(in *input.Input) Result[string] {
		strLen := len(str)
		startPos := in.Position

		end := len(in.Code) - strLen

		for in.Position <= end {
			p := in.Position

			if in.Code[p:p+strLen] == str {
				in.SetPosition(p + strLen)
				return Result[string]{Value: in.Code[startPos:p]}
			}

			in.Advance()
		}

		in.SetPosition(startPos)
		return Result[string]{Failed: true}
	}
}

func Many[T any](parser Parser[T]) Parser[[]T] {
	return func(in *input.Input) Result[[]T] {
		var results []T

		for !in.End() {
			result := parser(in)

			if result.Failed {
				break
			} else {
				results = append(results, result.Value)
			}
		}

		return Result[[]T]{
			Value: results,
		}
	}
}

func Many1[T any](parser Parser[T]) Parser[[]T] {
	return func(in *input.Input) Result[[]T] {
		var results []T

		for !in.End() {
			result := parser(in)

			if result.Failed {
				break
			} else {
				results = append(results, result.Value)
			}
		}

		return Result[[]T]{
			Failed: len(results) == 0,
			Value:  results,
		}
	}
}
