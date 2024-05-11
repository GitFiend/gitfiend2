package parser

import (
	"fmt"
)

type Result[T any] struct {
	Failed bool
	Value  T
}

type Parser[T any] func(i *Input) Result[T]
type Parser2[T any] func(i *Input) (T, bool)

func Parse[T any](parser Parser[T], text string) Result[T] {
	return parseInner(parser, text, true)
}

func ParsePart[T any](parser Parser[T], text string) Result[T] {
	return parseInner(parser, text, false)
}

// parseInner
// TODO: We shouldn't be printing for every failure? Some of our tests expect failure and it's annoying.
func parseInner[T any](parser Parser[T], text string, mustParseAll bool) Result[T] {
	in := NewInput(text)
	result := parser(&in)

	if mustParseAll && !in.End() {
		message := fmt.Sprintf(
			`
PARSE FAILURE AT POSITION %d:
  SUCCESSFULLY PARSED:
  "%s"

  FAILED AT:
  "%s"
`,
			in.AttemptedPosition,
			in.SuccessfullyParsed(),
			in.UnParsed(),
		)

		fmt.Println(message)
	}

	return result
}

// Map See tests for how to use this.
func Map[T any, U any](parser Parser[T], f func(result T) U) Parser[U] {
	return func(i *Input) Result[U] {
		res := parser(i)

		if !res.Failed {
			return Result[U]{
				Value: f(res.Value),
			}
		}
		return Result[U]{Failed: true}
	}
}
