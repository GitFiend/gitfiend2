package parser

type Input struct {
	Code              []rune
	Len               int
	Position          int
	AttemptedPosition int
}

func NewInput(text string) Input {
	code := []rune(text)
	return Input{
		Code: code,
		Len:  len(code),
	}
}

func (in *Input) Advance() {
	in.SetPosition(in.Position + 1)
}

func (in *Input) End() bool {
	return in.Position >= len(in.Code)
}

func (in *Input) AdvanceBy(pos int) {
	in.SetPosition(in.Position + pos)
}

func (in *Input) SetPosition(pos int) {
	if pos > in.AttemptedPosition {
		in.AttemptedPosition = pos
	}
	in.Position = pos
}

func (in *Input) NextRune() rune {
	return in.Code[in.Position]
}

func (in *Input) Rest(chunkLength int) string {
	return string(in.Code[in.Position:min(len(in.Code), in.Position+chunkLength)])
}

func (in *Input) SuccessfullyParsed() string {
	return string(in.Code[0:in.AttemptedPosition])
}

func (in *Input) UnParsed() string {
	return string(in.Code[in.AttemptedPosition:])
}
