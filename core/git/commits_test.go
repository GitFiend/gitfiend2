package git

import (
	"fmt"
	"gitfiend2/core"
	"gitfiend2/core/parser"
	"testing"
)

func TestPCommitRow(t *testing.T) {
	text := fmt.Sprintf(
		`Toby, sugto555@gmail.com, 1648863350 +1300, dd5733ad96082f0f33164bd1e2d72f7540bf7d9f, 2e8966986f620f491c34e6243a546d85dd2322e0, Write commit row parser. Added necessary new git types. %s,  (HEAD -> refs/heads/master, refs/remotes/origin/master)`,
		End,
	)

	t.Run(
		text, func(t *testing.T) {
			_, ok := parser.Parse(PCommitRow, text)

			if !ok {
				t.Error("Expected success")
			}
		},
	)
}

func TestLoadCommits(t *testing.T) {
	//dir, _ := os.Getwd()

	defer core.Elapsed("LoadCommits")()

	res := LoadCommits(RunOpts{RepoPath: `/home/toby/Repos/vscode`}, 5000)

	println(len(res))
}

func TestPDate(t *testing.T) {
	res, ok := parser.Parse(pDate, "1243 23")

	if !ok {
		t.Error("Expected success")
	}
	if res.Ms != 1243000 {
		t.Error("Expected 1243000, got ", res.Ms)
	}
	if res.Adjustment != 23 {
		t.Error("Expected 23, got ", res.Adjustment)
	}
}

func TestPParents(t *testing.T) {
	t.Run(
		"2 parents", func(t *testing.T) {
			h1 := "914aca5d9be2674304564e83efdcba92267dd7f9"
			h2 := "505586ea2ec4431a462d9e37cff7750923b199f0"
			var text = h1 + " " + h2

			res, ok := parser.Parse(PParents, text)

			if !ok {
				t.Error(`Failed to parse ` + text)
			}
			if res[0] != h1 {
				t.Error(`Failed to get ` + h1)
			}
			if res[1] != h2 {
				t.Error(`Failed to get ` + h2)
			}
		},
	)

	t.Run(
		"no parents", func(t *testing.T) {

			_, ok := parser.Parse(PParents, "")

			if !ok {
				t.Error(`Expected success when there's no parent hashes to parse.'`)
			}
		},
	)
}

func TestPMessage(t *testing.T) {
	t.Run(
		`Parse exmaple message 1`, func(t *testing.T) {
			res, ok := parser.Parse(PMessage, `fasdf *\nasdf `+End+` asdf`)

			if !ok {
				t.Error(`Expected success`)
			}
			if res != `fasdf *\nasdf ` {
				t.Error(`Expected Value "fasdf *\nasdf "`)
			}
		},
	)

	t.Run(
		`Parse realistic message`, func(t *testing.T) {
			text := fmt.Sprintf(`Write commit row parser. Added necessary new git types. %s`, End)

			_, ok := parser.Parse(PMessage, text)

			if !ok {
				t.Error()
			}
		},
	)
}
