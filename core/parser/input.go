package parser

import (
	"gitfiend2/core"
)

type Input struct {
	Code              string
	Position          int
	AttemptedPosition int
}

func (in *Input) Advance() {
	in.SetPosition(in.Position + 1)
}

func (in *Input) End() bool {
	return in.Position >= len(in.Code)
}

//func (in *Input) EndOfInput() bool {
//	length := len(in.Code)
//
//	return in.Position == length && in.AttemptedPosition == length
//}

func (in *Input) AdvanceBy(pos int) {
	in.SetPosition(in.Position + pos)
}

func (in *Input) SetPosition(pos int) {
	if pos > in.AttemptedPosition {
		in.AttemptedPosition = pos
	}
	in.Position = pos
}

func (in *Input) NextChar() rune {
	return rune(in.Code[in.Position])
}

func (in *Input) Rest(chunkLength int) string {
	return in.Code[in.Position:core.Min(len(in.Code), in.Position+chunkLength)]
}

func (in *Input) SuccessfullyParsed() string {
	return in.Code[0:in.AttemptedPosition]
}

func (in *Input) UnParsed() string {
	return in.Code[in.AttemptedPosition:]
}
