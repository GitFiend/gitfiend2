package git

import . "gitfiend2/core/parser"

type Heading [2]string

type Row [2]string

type Other struct {
	Kind  OtherKind
	Value string
}

type OtherKind string

const (
	Comment OtherKind = "comment"
	Unknown           = "unknown"
)

var pHeading1 = Map(And3(Rune('['), AnyWord, Rune(']')),
	func(result And3Result[rune, string, rune]) Heading {
		return Heading{
			result.R2,
			"",
		}
	})

var pHeading2 = Map(And5(Rune('['), AnyWord, Ws, StringLiteral, Rune(']')),
	func(res And5Result[rune, string, string, string, rune]) Heading {
		return Heading{res.R2, res.R4}
	})

var pHeading = Or(pHeading1, pHeading2)

var pRow = Map(And6(Ws, AnyWord, Ws, Rune('='), Ws, UntilLineEnd),
	func(res And6Result[string, string, string, rune, string, string]) Row {
		return Row{res.R2, res.R6}
	})
