package git

import (
	. "gitfiend2/core/parser"
	"strings"
)

func ParseConfig(text string) ([]Row, bool) {
	return ParseAll(pConfig, text)
}

func MakeConfigLog(text string) (string, bool) {
	res, ok := ParseAll(pConfig, text)

	if ok {
		var log string
		for _, row := range res {
			if row.IsData() {
				log += row.String() + "\n"
			}
		}
		return strings.Trim(log, "\n "), true
	}
	return "", false
}

type Row interface {
	String() string
	IsData() bool
}

type Section struct {
	Heading Heading
	Rows    []Row
}

func (s Section) String() string {
	var res string
	heading := s.Heading.String()

	for _, row := range s.Rows {
		if row.IsData() {
			res += heading + "." + row.String() + "\n"
		}
	}
	return strings.Trim(res, "\n ")
}
func (s Section) IsData() bool {
	return true
}

func (s Section) Entries() [][2]string {
	heading := s.Heading.String()
	var entries [][2]string

	for _, row := range s.Rows {
		switch r := row.(type) {
		case DataRow:
			entries = append(entries, [2]string{heading + "." + r[0], r[1]})
			break
		}
	}
	return entries
}

type Heading [2]string

func (h Heading) String() string {
	if h[1] == "" {
		return h[0]
	} else {
		return h[0] + "." + h[1]
	}
}
func (h Heading) Key() string {
	return h[0]
}
func (h Heading) Value() string {
	return h[1]
}

type DataRow [2]string

func (r DataRow) String() string {
	return r[0] + "=" + r[1]
}
func (r DataRow) IsData() bool {
	if r[0] == "" || r[1] == "" {
		return false
	}
	return true
}
func (r DataRow) Key() string {
	return r[0]
}
func (r DataRow) Value() string {
	return r[1]
}

type Unknown string
type Comment string

func (c Comment) String() string {
	return string(c)
}
func (c Comment) IsData() bool {
	return false
}
func (o Unknown) String() string {
	return string(o)
}
func (o Unknown) IsData() bool {
	return false
}

var pConfig = Many(Or(pSection, pOther))

var pHeading = Or(pHeading1, pHeading2)

var pHeading1 = Map(
	And3(Rune('['), AnyWord, Rune(']')),
	func(result And3Result[rune, string, rune]) Heading {
		return Heading{result.R2}
	},
)

var pHeading2 = Map(
	And5(Rune('['), AnyWord, Ws, StringLiteral, Rune(']')),
	func(res And5Result[rune, string, string, string, rune]) Heading {
		return Heading{res.R2, res.R4}
	},
)

var pRow = Map(
	And6(Ws, AnyWord, Ws, Rune('='), Ws, UntilLineEnd),
	func(res And6Result[string, string, string, rune, string, string]) Row {
		return DataRow{res.R2, res.R6}
	},
)

var pSection = Map(
	And2(pHeading, Many(Or(pRow, pComment, pUnknown))),
	func(res And2Result[Heading, []Row]) Row {
		return Section{res.R1, res.R2}
	},
)

var pComment = Map(
	And3(Ws, Or(Rune(';'), Rune('#')), UntilLineEnd),
	func(res And3Result[string, rune, string]) Row {
		return Comment(res.R3)
	},
)

var pUnknown = Map(
	And3(Not(pHeading), Not(pRow), UntilLineEnd),
	func(res And3Result[bool, bool, string]) Row {
		return Unknown(res.R3)
	},
)

var pOther = Or(pComment, pUnknown)
