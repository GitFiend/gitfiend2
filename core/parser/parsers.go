package parser

import (
	"regexp"
	"slices"
	"unicode"
)

func Rune(c rune) Parser[rune] {
	return func(in *Input) (rune, bool) {
		if !in.End() {
			n := in.NextRune()
			if n == c {
				in.Advance()
				return n, true
			}
		}
		return *new(rune), false
	}
}

func Word(word string) Parser[string] {
	return func(in *Input) (string, bool) {
		p := in.Position

		for _, c := range word {
			if !in.End() && in.NextRune() == c {
				in.Advance()
			} else {
				in.SetPosition(p)
				return "", false
			}
		}

		return word, true
	}
}

// Not Succeeds if it fails to parse. Doesn't change position.
func Not[T any](parser Parser[T]) Parser[bool] {
	return func(in *Input) (bool, bool) {
		p := in.Position
		_, ok := parser(in)

		if ok {
			in.SetPosition(p)
			return true, false
		}

		return false, true
	}
}

// Regex This is really slow, so only use it if that doesn't matter.
func Regex(re *regexp.Regexp) Parser[string] {
	return func(i *Input) (string, bool) {
		code := i.Code[i.Position:]
		match := re.FindStringIndex(string(code))

		if match[0] == 0 {
			start := i.Position
			i.AdvanceBy(match[1])

			return string(i.Code[start : start+match[1]]), true
		}

		return "", false
	}
}

func OptionalWhiteSpace[T any]() Parser[T] {
	return func(in *Input) (T, bool) {
		for !in.End() && unicode.IsSpace(in.NextRune()) {
			in.Advance()
		}

		return *new(T), true
	}
}

func Optional[T any](parser Parser[T]) Parser[T] {
	return func(in *Input) (T, bool) {
		res, ok := parser(in)
		if ok {
			return res, true
		}
		return *new(T), true
	}
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
				break
			}
			results = append(results, result)

			if in.End() {
				break
			}

			_, ok = separator(in)
			if !ok {
				break
			}
		}

		return results, true
	}
}

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
	}
}

// UntilParser
// Parses until parser is found or the end of input. Always succeeds.
// All text is consumed, but end parser result is not included (TODO: Check this)
func UntilParser[T any](parser Parser[T]) Parser[string] {
	return func(in *Input) (string, bool) {
		startPos := in.Position
		currentPos := startPos

		for !in.End() {
			currentPos = in.Position
			_, ok := parser(in)

			if ok {
				break
			}
			in.Advance()
		}

		if startPos == currentPos {
			return "", true
		}
		return string(in.Code[startPos:currentPos]), true
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
	}
}

func TakeRuneWhile(f func(r rune) bool) Parser[string] {
	return func(in *Input) (string, bool) {
		startPos := in.Position

		for !in.End() && f(in.NextRune()) {
			in.Advance()
		}

		if startPos == in.Position {
			return "", false
		}
		return string(in.Code[startPos:in.Position]), true
	}
}

func ConditionalRune(f func(r rune) bool) Parser[rune] {
	return func(i *Input) (rune, bool) {
		r := i.NextRune()

		if f(r) {
			i.Advance()
			return r, true
		}
		return *new(rune), false
	}
}
