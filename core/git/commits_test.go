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
		"Firstname Lastname; somewords123@gmail.com; 1648863350 +1300; dd5733ad96082f0f33164bd1e2d72f7540bf7d9f;"+
			" 2e8966986f620f491c34e6243a546d85dd2322e0; Write commit row parser. Added necessary new git types. %s;"+
			"  (HEAD -> refs/heads/master, refs/remotes/origin/master)",
		End,
	)

	c, ok := parser.ParseAll(PCommitRow, text)

	assert.True(t, ok)
	assert.Equal(t, 1648863350000, c.Date.Ms)
	assert.Equal(t, "Firstname Lastname", c.Author)
}

func TestLoadCommits(t *testing.T) {
	defer shared.Elapsed("LoadCommits")()

	home, _ := os.UserHomeDir()
	res := LoadCommits(path.Join(home, "Repos", "vscode"), 5000)

	println(len(res))
}

func TestPDate(t *testing.T) {
	res, ok := parser.ParseAll(pDate, "1243 23")

	assert.True(t, ok)
	assert.Equal(t, 1243000, res.Ms)
	assert.Equal(t, 23, res.Adjustment)
}

func TestSliceExpectations(t *testing.T) {
	ids := []string{"1", "2", "3"}
	trimmed := ids[:1]

	assert.Equal(t, []string{"1"}, trimmed)
}

func TestPParents(t *testing.T) {
	t.Run(
		"2 parents", func(t *testing.T) {
			h1 := "914aca5d9be2674304564e83efdcba92267dd7f9"
			h2 := "505586ea2ec4431a462d9e37cff7750923b199f0"
			text := h1 + " " + h2

			res, ok := parser.ParseAll(PParents, text)
			assert.True(t, ok)
			assert.Equal(t, h1, res[0])
			assert.Equal(t, h2, res[1])
		},
	)

	t.Run(
		"no parents", func(t *testing.T) {
			_, ok := parser.ParseAll(PParents, "")
			assert.True(t, ok)
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
			assert.True(t, ok)
		},
	)
}
