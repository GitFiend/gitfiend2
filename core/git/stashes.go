package git

import (
	"fmt"
	"gitfiend2/core/parser"
	"strings"
)

func LoadStashes(repoPath string) []CommitInfo {
	res, err := RunGit(RunOpts{RepoPath: repoPath, Args: []string{
		"reflog",
		"show",
		"stash",
		"-z",
		"--decorate=full",
		prettyFormatted,
		"--date=raw",
	}})
	if err != nil {
		return nil
	}

	commits, ok := parser.ParseAll(PCommits, res.Stdout)
	if ok {
		for i, c := range commits {
			c.StashId = fmt.Sprintf("refs/stash@{%d}", i)
			c.IsMerge = false
			c.Ref = nil

			if len(c.ParentIds) > 1 {
				c.ParentIds = c.ParentIds[:1]
			}
			c.Message = tidyCommitMessage(c.Message)
		}
		return commits
	}
	return nil
}

func tidyCommitMessage(message string) string {
	parts := strings.Split(message, ":")
	if len(parts) > 1 {
		m := parts[1]
		return strings.Replace(m, "WIP", "Stash", 1)
	}
	return "Stash"
}
