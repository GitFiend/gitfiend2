package git

import (
	"os"
	"testing"
)

func TestRunGit(t *testing.T) {
	path, err := os.Getwd()

	if err != nil {
		t.Error("os.Getwd() failed")
	}

	runResult, err := RunGit(RunOpts{
		Args:     []string{"--version"},
		RepoPath: path,
	})

	if err != nil || runResult.Stdout == "" {
		t.Error("Expected result test, got nothing.")
	}
}
