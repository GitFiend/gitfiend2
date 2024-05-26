package git

import (
	"fmt"
	"gitfiend2/core/parser"
	"gitfiend2/core/shared"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestPCommitRow(t *testing.T) {
	text := fmt.Sprintf(
		"Toby; sugto555@gmail.com; 1648863350 +1300; dd5733ad96082f0f33164bd1e2d72f7540bf7d9f;"+
			" 2e8966986f620f491c34e6243a546d85dd2322e0; Write commit row parser. Added necessary new git types. %s;"+
			"  (HEAD -> refs/heads/master, refs/remotes/origin/master)",
		End,
	)

	t.Run(
		text, func(t *testing.T) {
			_, ok := parser.ParseAll(PCommitRow, text)

			if !ok {
				t.Error("Expected success")
			}
		},
	)
}

func TestLoadCommits(t *testing.T) {
	defer shared.Elapsed("LoadCommits")()

	home, _ := os.UserHomeDir()
	res := LoadCommits(RunOpts{RepoPath: path.Join(home, "Repos", "vscode")}, 5000)

	println(len(res))
}

func TestPDate(t *testing.T) {
	p := parser.New(pDate, "1243 23")
	res, ok := p.Run()

	assert.True(t, p.Finished())

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

			res, ok := parser.ParseAll(PParents, text)

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

			_, ok := parser.ParseAll(PParents, "")

			if !ok {
				t.Error(`Expected success when there's no parent hashes to parse.'`)
			}
		},
	)
}

func TestPMessage(t *testing.T) {
	t.Run(
		`ParseAll exmaple message 1`, func(t *testing.T) {
			res, ok := parser.ParsePart(PMessage, `fasdf *\nasdf `+End+` asdf`)

			assert.True(t, ok)
			assert.Equal(t, `fasdf *\nasdf `, res)
		},
	)

	t.Run(
		`ParseAll realistic message`, func(t *testing.T) {
			text := fmt.Sprintf(`Write commit row parser. Added necessary new git types. %s`, End)

			_, ok := parser.ParseAll(PMessage, text)

			if !ok {
				t.Error()
			}
		},
	)
}
