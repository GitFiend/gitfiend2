package git

import (
	"fmt"
	. "gitfiend2/core/parser"
	"gitfiend2/core/shared"
	"os"
	"strconv"
)

const End = "4a41380f-a4e8-4251-9ca2-bf55186ed32a"
const SepRune = ';'

const commit0Id = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"

/**
 * %an author name
 * %ae author email
 * %ad date
 * %H hash
 * %P parent hashes
 * %B message
 * %d remotes
 */
const prettyFormatted = `--pretty=format:%an; %ae; %ad; %H; %P; %B` + End + `; %d`

var pGroup = TakeRuneWhile(
	func(r rune) bool {
		return r != SepRune
	},
)

var pSep = Map(
	And3(Ws, Rune(SepRune), Ws), func(result And3Result[string, rune, string]) rune {
		return SepRune
	},
)

var pEmail = Or(pGroup, Ws)

var pDate = Map(
	And3(UInt, Ws, SignedInt), func(result And3Result[string, string, string]) DateResult {
		micro, _ := strconv.Atoi(result.R1)
		adjustment, _ := strconv.Atoi(result.R3)

		return DateResult{Ms: micro * 1000, Adjustment: adjustment}
	},
)

var PParents = RepParserSep(AnyWord, Ws)

var PMessage = UntilString(End)

var PIdList = RepParserSep(AnyWord, UntilLineEnd)

var PCommitRow = Map(
	And14(
		/* R1 */ pGroup, // author
		/* R2 */ pSep,
		/* R3 */ pEmail, // email
		/* R4 */ pSep,
		/* R5 */ pDate,
		/* R6 */ pSep,
		/* R7 */ pGroup, // commit id
		/* R8 */ pSep,
		/* R9 */ PParents,
		/* R10 */ pSep,
		/* R11 */ PMessage,
		/* R12 */ pSep,
		/* R13 */ POptionalRefs,
		/* R14 */ Ws,
	),
	func(
		result And14Result[
			string,
			rune,
			string,
			rune,
			DateResult,
			rune,
			string,
			rune,
			[]string,
			rune,
			string,
			rune,
			[]refInfoPart,
			string,
		],
	) CommitInfo {
		c := CommitInfo{
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

func LoadCommits(repoPath string, num int) []CommitInfo {
	fmt.Println(os.Environ())

	options := RunOpts{
		RepoPath: repoPath,
		Args: []string{
			"log",
			"--branches",
			"--tags",
			"--remotes",
			"--decorate=full",
			prettyFormatted,
			"-n" + strconv.Itoa(num),
			"--date=raw",
		},
	}

	textResult, err := RunGit(options)
	if err != nil {
		return nil
	}

	defer shared.Elapsed("ParseAll commits")()
	res, _ := ParseAll(PCommits, textResult.Stdout)

	return res
}
