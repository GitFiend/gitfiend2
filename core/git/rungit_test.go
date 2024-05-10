package gitTypes

import (
	"os"
	"testing"
)

func TestRunGit(t *testing.T) {
	path, err := os.Getwd()

	if err != nil {
		t.Error("os.Getwd() failed")
	}

	res := RunGit(GitOptions{
		Args:     []string{"--version"},
		RepoPath: path,
	})

	if len(res) == 0 {
		t.Error("Expected result test, got nothing.")
	}
}
