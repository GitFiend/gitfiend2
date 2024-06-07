package store

import (
	"slices"
)

var repoPaths []RepoPath
var commitsAndRefs CommitsAndRefs

func SetRepoPaths(repos []RepoPath) {
	repoPaths = repos
}

func GetRepoPath(repoPath string) (RepoPath, bool) {
	i := slices.IndexFunc(repoPaths, func(p RepoPath) bool {
		return p.Path == repoPath
	})

	if i >= 0 {
		return repoPaths[i], true
	}
	return RepoPath{}, false
}

func SetCommitsAndRefs(c CommitsAndRefs) {
	commitsAndRefs = c
}
