package runGit

import (
	"os/exec"
)

type GitOptions struct {
	Args      []string
	RepoPath  string
	Timeout   bool
	ShowError bool
}

func RunGit(options GitOptions) string {
	cmd := exec.Command("git", options.Args...)

	cmd.Dir = options.RepoPath
	stdoutStderr, _ := cmd.CombinedOutput()

	return string(stdoutStderr)
}
