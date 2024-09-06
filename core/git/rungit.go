package git

import (
	"bytes"
	"os/exec"
)

type RunOpts struct {
	Args     []string
	RepoPath string

	// TODO
	Timeout   bool
	ShowError bool
}

type RunResult struct {
	Stdout string
	Stderr string
}

func RunGit(options RunOpts) (RunResult, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("git", options.Args...)
	cmd.Dir = options.RepoPath

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return RunResult{}, err
	}

	return RunResult{Stdout: stdout.String(), Stderr: stderr.String()}, nil
}
