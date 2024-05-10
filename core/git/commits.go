package gitTypes

import (
	"gitfiend2/core"
	. "gitfiend2/core/parser"
	"os"
	"regexp"
	"strconv"
)

var End = core.PseudoUuid()

/**
 * %an author name
 * %ae author email
 * %ad date
 * %H hash
 * %P parent hashes
 * %B message
 * %d remotes
 */
var prettyFormatted = `--pretty=format:%an, %ae, %ad, %H, %P, %B` + End + `, %d`

var pGroup = Regex(regexp.MustCompile(`[^,]+`))

var pSep = And3(Ws, Char(','), Ws)

type SepResult = And3Result[string, rune, string]

var pDate = Map(
	And3(Uint, Ws, Int), func(result And3Result[string, string, string]) DateResult {
		micro, _ := strconv.Atoi(result.A)
		adjustment, _ := strconv.Atoi(result.C)

		return DateResult{Ms: micro * 1000, Adjustment: adjustment}
	},
)

var PParents = RepParserSep(AnyWord, Ws)

var PMessage = Until(End)

var PCommitRow = Map(
	And14(
		/* A */ pGroup, // author
		/* B */ pSep,
		/* C */ Or(pGroup, Ws), // email
		/* D */ pSep,
		/* E */ pDate,
		/* F */ pSep,
		/* G */ pGroup, // commit id
		/* H */ pSep,
		/* I */ PParents,
		/* J */ pSep,
		/* K */ PMessage,
		/* L */ pSep,
		/* M */ POptionalRefs,
		/* N */ Ws,
	),
	func(
		result And14Result[
			string,
			SepResult,
			string,
			SepResult,
			DateResult,
			SepResult,
			string,
			SepResult,
			[]string,
			SepResult,
			string,
			SepResult,
			[]RefInfoPart,
			string,
		],
	) Commit {
		c := Commit{
			Author:    result.A,
			Email:     result.C,
			Date:      result.E,
			Id:        result.G,
			Index:     0,
			ParentIds: result.I,
			IsMerge:   len(result.I) == 2,
			Message:   result.K,
		}

		if len(result.M) > 0 {
			for _, info := range result.M {
				ref := MakeRefInfo(info)
				ref.CommitId = c.Id
				ref.Time = c.Date.Ms
				c.Ref = append(c.Ref, ref)
			}
		}

		return c
	},
)

var PCommits = Many(PCommitRow)

// LoadCommits
// Intentional copy of options, so we can modify it.
func LoadCommits(options GitOptions, num uint) []Commit {
	print(os.Environ())

	options.Args = []string{
		"log",
		"--branches",
		"--tags",
		"--remotes",
		"--decorate=full",
		prettyFormatted,
		"-n" + strconv.Itoa(int(num)),
		"--date=raw",
	}

	textResult := RunGit(options)

	if len(textResult) > 0 {
		defer core.Elapsed("Parse commits")()

		res := Parse(PCommits, textResult)

		return res.Value
	}

	return nil
}
