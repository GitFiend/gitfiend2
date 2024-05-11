package parser

import (
	"regexp"
	"slices"
)

func Char(c rune) Parser[rune] {
	return func(in *Input) (rune, bool) {
		if !in.End() {
			n := in.NextChar()

			if n == c {
				in.Advance()
				return n, true
				//return Result[rune]{Value: n}
			}
		}

		return *new(rune), false
		//return Result[rune]{Failed: true}
	}
}

func Word(word string) Parser[string] {
	return func(in *Input) (string, bool) {
		p := in.Position

		for _, c := range word {
			if !in.End() && in.NextChar() == c {
				in.Advance()
			} else {
				in.SetPosition(p)
				return "", false
				//return Result[string]{Failed: true}
			}
		}

		return word, true
		//return Result[string]{Value: word}
	}
}

// Regex
// var re *regexp.Regexp = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
func Regex(re *regexp.Regexp) Parser[string] {
	return func(i *Input) (string, bool) {
		code := i.Code[i.Position:]
		match := re.FindStringIndex(string(code))

		if match[0] == 0 {
			start := i.Position
			i.AdvanceBy(match[1])

			//return Result[string]{Value: string(i.Code[start : start+match[1]])}
			return string(i.Code[start : start+match[1]]), true
		}

		return "", false
		//return Result[string]{Failed: true}
	}
}

func OptionalWhiteSpace[T any]() Parser[T] {
	return func(in *Input) (T, bool) {
		for !in.End() && IsWhiteSpace(in.NextChar()) {
			in.Advance()
		}

		return *new(T), true
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

// RepParserSep Doesn't require a result, always succeeds.
func RepParserSep[T any, U any](parser Parser[T], separator Parser[U]) Parser[[]T] {
	return func(in *Input) ([]T, bool) {
		var results []T

		for !in.End() {
			result, ok := parser(in)
			if !ok {
				//if len(results) == 0 {
				//	//return Result[[]T]{Failed: true}
				//	return Result[[]T]{}
				//}

				break
			}
			results = append(results, result)

			if in.End() {
				break
				//return Result[[]T]{Value: results}
			}

			_, ok = separator(in)
			if !ok {
				break
			}
		}

		return results, true
	}
}

//// Until
//// Input is consumed including str, but str is not included in the result.
//func Until(str string) Parser[string] {
//	return func(in *Input) Result[string] {
//		strLen := len(str)
//		startPos := in.Position
//
//		end := len(in.Code) - strLen
//
//		for in.Position <= end {
//			p := in.Position
//
//			if in.Code[p:p+strLen] == str {
//				in.SetPosition(p + strLen)
//				return Result[string]{Value: in.Code[startPos:p]}
//			}
//
//			in.Advance()
//		}
//
//		in.SetPosition(startPos)
//		return Result[string]{Failed: true}
//	}
//}

// UntilString
// Input is consumed including str, but str is not included in the result.
func UntilString(str string) Parser[string] {
	return func(in *Input) (string, bool) {
		runes := []rune(str)
		strLen := len(runes)
		startPos := in.Position
		end := in.Len - strLen

		for in.Position <= end {
			p := in.Position

			if slices.Equal(in.Code[p:p+strLen], runes) {
				in.SetPosition(p + strLen)
				return string(in.Code[startPos:p]), true
			}

			in.Advance()
		}

		in.SetPosition(startPos)
		return "", false
		//return Result[string]{Failed: true}
	}
}

func Many[T any](parser Parser[T]) Parser[[]T] {
	return func(in *Input) ([]T, bool) {
		var results []T

		for !in.End() {
			result, ok := parser(in)

			if !ok {
				break
			} else {
				results = append(results, result)
			}
		}

		return results, true

		//return Result[[]T]{
		//	Value: results,
		//}
	}
}

func Many1[T any](parser Parser[T]) Parser[[]T] {
	return func(in *Input) ([]T, bool) {
		var results []T

		for !in.End() {
			result, ok := parser(in)

			if !ok {
				break
			} else {
				results = append(results, result)
			}
		}

		return results, len(results) > 0
		//return Result[[]T]{
		//	Failed: len(results) == 0,
		//	Value:  results,
		//}
	}
}

//func TakeCharWhile(f func(r rune) bool) Parser[string] {
//	s := "asdf"
//
//	return func(in *Input) Result[string] {
//		for !in.End() {
//			if f(in.NextChar()) {
//				in.Advance()
//			} else {
//				break
//			}
//		}
//	}
//}
