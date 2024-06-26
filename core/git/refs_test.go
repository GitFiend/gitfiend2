package git

import (
	"gitfiend2/core/parser"
	"testing"
)

func TestGetRefLocation(t *testing.T) {
	origin := getRefLocation([]string{"refs", "heads", "commit-list-experiments"})

	if origin != Local {
		t.Error("Expected local ref")
	}
}

func TestGetRefShortName(t *testing.T) {
	n1 := getShortName([]string{"refs", "heads", "feature", "dialogs"})

	if n1 != "feature/dialogs" {
		t.Error(`Expected "feature/dialogs"`)
	}

	n2 := getShortName([]string{"refs", "remotes", "origin", "git-lib"})

	if n2 != "git-lib" {
		t.Error(`Expected "git-lib"`)
	}
}

func TestGetRemoteName(t *testing.T) {
	if getRemoteName([]string{"refs", "remotes", "origin", "git-lib"}) != "origin" {
		t.Error(`Expected "origin"`)
	}
	if getRemoteName([]string{"refs", "heads", "feature", "dialogs"}) != "" {
		t.Error(`Expected no remote for ref`)
	}
	if getRemoteName([]string{"refs", "tags", "hello"}) != "" {
		t.Error(`Expected no remote`)
	}
}

func TestPRefName(t *testing.T) {
	_, ok := parser.ParseAll(PRefName, "refs/heads/git-lib")

	if !ok {
		t.Error(`Expected parser success`)
	}
}

func TestPTagRef(t *testing.T) {
	_, ok := parser.ParseAll(PTagRef, `tag: refs/tags/v0.11.2`)

	if !ok {
		t.Error(`Expected parser success`)
	}
}

func TestPHeadRef(t *testing.T) {
	res, ok := parser.ParseAll(PHeadRef, `HEAD -> refs/heads/master`)

	if !ok {
		t.Error(`Expected parse success`)
	}
	if res.Id != `refs/heads/master` {
		t.Error(`Expected "refs/heads/master", got ` + res.Id)
	}
}

func TestPCommitRefs(t *testing.T) {
	a := `(HEAD -> refs/heads/master, refs/remotes/origin/master, refs/remotes/origin/HEAD)`

	res, ok := parser.ParseAll(PCommitRefs, a)

	if !ok {
		t.Error(`Expected success`)
	}
	if len(res) != 3 {
		t.Error(`Expected 3 refs`)
	}
	if res[1].Id != `refs/remotes/origin/master` {
		t.Error(`Expected 2nd ref id match`)
	}
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
