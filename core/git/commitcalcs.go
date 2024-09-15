package git

import (
	"fmt"
	"strconv"
	"strings"
)

func countCommitsBetweenFallback(repoPath, commitId1, commitId2 string) int {
	if commitId1 == commitId2 {
		return 0
	}

	res, err := RunGit(RunOpts{RepoPath: repoPath, Args: []string{
		"rev-list",
		fmt.Sprintf("%v..%v", commitId1, commitId2),
		"--count",
	}})

	if err == nil {
		n, err := strconv.Atoi(strings.TrimSpace(res.Stdout))
		if err == nil {
			return n
		}
	}
	return 0
}
