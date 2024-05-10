package parser

import (
	"fmt"
	"gitfiend2/core/input"
)

type Result[T any] struct {
	Failed bool
	Value  T
}

type Parser[T any] func(i *input.Input) Result[T]

func Parse[T any](parser Parser[T], text string) Result[T] {
	return ParseInner(parser, text, true)
}

func ParsePart[T any](parser Parser[T], text string) Result[T] {
	return ParseInner(parser, text, false)
}

// TODO: We shouldn't be printing for every failure? Some of our tests expect failure and it's annoying.
func ParseInner[T any](parser Parser[T], text string, mustParseAll bool) Result[T] {
	in := input.Input{Code: text}
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
	return func(i *input.Input) Result[U] {
		res := parser(i)

		if !res.Failed {
			return Result[U]{
				Value: f(res.Value),
			}
		}
		return Result[U]{Failed: true}
	}
}
