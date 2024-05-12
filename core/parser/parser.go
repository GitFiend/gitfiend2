package parser

import "fmt"

type Parser[T any] func(i *Input) (T, bool)

func Parse[T any](parser Parser[T], text string) (T, bool) {
	return parseInner(parser, text, true)
}

func ParsePart[T any](parser Parser[T], text string) (T, bool) {
	return parseInner(parser, text, false)
}

type ParseOpts struct {
}

type ParseInstance[T any] struct {
	text     string
	input    *Input
	parser   Parser[T]
	parseAll bool
}

func RunParse[T any](parser Parser[T], text string, parseAll bool) *ParseInstance[T] {
	in := NewInput(text)

	return &ParseInstance[T]{
		text:     text,
		input:    &in,
		parser:   parser,
		parseAll: parseAll,
	}
}

func (p *ParseInstance[T]) Run() (T, bool) {
	in := p.input
	result, ok := p.parser(in)

	if p.parseAll && !in.End() {
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

	return result, ok
}

// parseInner
// TODO: We shouldn't be printing for every failure? Some of our tests expect failure and it's annoying.
func parseInner[T any](parser Parser[T], text string, mustParseAll bool) (T, bool) {
	in := NewInput(text)
	result, ok := parser(&in)

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

	return result, ok
}

// Map See tests for how to use this.
func Map[T any, U any](parser Parser[T], f func(result T) U) Parser[U] {
	return func(i *Input) (U, bool) {
		res, ok := parser(i)

		if ok {
			return f(res), true
		}
		return *new(U), false
	}
}
