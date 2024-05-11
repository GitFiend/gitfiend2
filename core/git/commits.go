package git

import (
	"fmt"
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
		micro, _ := strconv.Atoi(result.R1)
		adjustment, _ := strconv.Atoi(result.R3)

		return DateResult{Ms: micro * 1000, Adjustment: adjustment}
	},
)

var PParents = RepParserSep(AnyWord, Ws)

var PMessage = UntilString(End)

var PCommitRow = Map(
	And14(
		/* A R1 */ pGroup, // author
		/* B R2 */ pSep,
		/* C R3 */ Or(pGroup, Ws), // email
		/* D R4 */ pSep,
		/* E R5 */ pDate,
		/* F R6 */ pSep,
		/* G R7 */ pGroup, // commit id
		/* H R8 */ pSep,
		/* I R9 */ PParents,
		/* J R10 */ pSep,
		/* K R11 */ PMessage,
		/* L R12 */ pSep,
		/* M R13 */ POptionalRefs,
		/* N R14 */ Ws,
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
			Author:    result.R1,
			Email:     result.R3,
			Date:      result.R5,
			Id:        result.R7,
			Index:     0,
			ParentIds: result.R9,
			IsMerge:   len(result.R9) == 2,
			Message:   result.R11,
		}

		if len(result.R13) > 0 {
			for _, info := range result.R13 {
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
func LoadCommits(options RunOpts, num uint) []Commit {
	fmt.Println(os.Environ())

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

	textResult, err := RunGit(options)
	if err != nil {
		return nil
	}

	defer core.Elapsed("Parse commits")()
	res, _ := Parse(PCommits, textResult.Stdout)

	return res
}
