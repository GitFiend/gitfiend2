package parser

import "fmt"

type Parser[T any] func(i *Input) (T, bool)

// ParseAll
// Useful for tests. No error info is kept.
func ParseAll[T any](parser Parser[T], text string) (T, bool) {
	p := New(parser, text)
	res, ok := p.Run()

	return res, ok && p.Finished()
}

// ParsePart
// Useful for tests. No error info is kept.
func ParsePart[T any](parser Parser[T], text string) (T, bool) {
	p := New(parser, text)
	return p.Run()
}

type ParseInstance[T any] struct {
	text   string
	input  *Input
	parser Parser[T]
}

func New[T any](parser Parser[T], text string) *ParseInstance[T] {
	in := NewInput(text)

	return &ParseInstance[T]{
		text:   text,
		input:  &in,
		parser: parser,
	}
}

func (p *ParseInstance[T]) Run() (T, bool) {
	return p.parser(p.input)
}

func (p *ParseInstance[T]) Finished() bool {
	return p.input.End()
}

func (p *ParseInstance[T]) GetErrorInfo() string {
	in := p.input

	if !in.End() {
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

		return message
	}
	return ""
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
