package git

import (
	"gitfiend2/core/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRefLocation(t *testing.T) {
	origin := getRefLocation([]string{"refs", "heads", "commit-list-experiments"})
	assert.Equal(t, "Local", string(origin))
}

func TestGetRefShortName(t *testing.T) {
	n1 := getShortName([]string{"refs", "heads", "feature", "dialogs"})
	assert.Equal(t, "feature/dialogs", n1)

	n2 := getShortName([]string{"refs", "remotes", "origin", "git-lib"})
	assert.Equal(t, "git-lib", n2)
}

func TestGetRemoteName(t *testing.T) {
	assert.Equal(t, "origin", getRemoteName([]string{"refs", "remotes", "origin", "git-lib"}))
	assert.Equal(t, "", getRemoteName([]string{"refs", "heads", "feature", "dialogs"}))
	assert.Equal(t, "", getRemoteName([]string{"refs", "heads", "feature", "dialogs"}))
}

func TestPRefName(t *testing.T) {
	_, ok := parser.ParseAll(PRefName, "refs/heads/git-lib")
	assert.True(t, ok)
}

func TestPTagRef(t *testing.T) {
	_, ok := parser.ParseAll(PTagRef, `tag: refs/tags/v0.11.2`)
	assert.True(t, ok)
}

func TestPHeadRef(t *testing.T) {
	res, ok := parser.ParseAll(PHeadRef, `HEAD -> refs/heads/master`)

	assert.True(t, ok)
	assert.Equal(t, `refs/heads/master`, res.Id)
	assert.True(t, res.Head)
}

func TestPCommitRefs(t *testing.T) {
	a := `(HEAD -> refs/heads/master, refs/remotes/origin/master, refs/remotes/origin/HEAD)`
	res, ok := parser.ParseAll(PCommitRefs, a)
	assert.True(t, ok)
	assert.Equal(t, 3, len(res))
	assert.Equal(t, `refs/remotes/origin/master`, res[1].Id)
}

func TestPOptionalRefs(t *testing.T) {
	t.Run(
		`ParseAll 2 refs`, func(t *testing.T) {
			a := `(HEAD -> refs/heads/master, refs/remotes/origin/master)`
			res, ok := parser.ParseAll(POptionalRefs, a)

			if !ok {
				t.Error(`Expected success`)
			}
			if len(res) != 2 {
				t.Error(`Expected 2 refs`)
			}
			if res[0].Id != `refs/heads/master` {
				t.Error(`Expected first ref id match`)
			}
		},
	)

	t.Run(
		`ParseAll 3 refs`, func(t *testing.T) {
			a := `(HEAD -> refs/heads/master, refs/remotes/origin/master, refs/remotes/origin/HEAD)`
			res, ok := parser.ParseAll(POptionalRefs, a)

			if !ok {
				t.Error(`Expected success`)
			}
			if len(res) != 3 {
				t.Error(`Expected 3 refs`)
			}
			if res[1].Id != `refs/remotes/origin/master` {
				t.Error(`Expected 2nd ref id match`)
			}
		},
	)

}
