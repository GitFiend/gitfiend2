package store

import (
	"gitfiend2/core"
	"slices"
)

var repoPaths []core.RepoPath
var commitsAndRefs CommitsAndRefs

func SetRepoPaths(repos []core.RepoPath) {
	repoPaths = repos
}

func GetRepoPath(repoPath string) (core.RepoPath, bool) {
	i := slices.IndexFunc(repoPaths, func(p core.RepoPath) bool {
		return p.Path == repoPath
	})

	if i >= 0 {
		return repoPaths[i], true
	}
	return core.RepoPath{}, false
}

func SetCommitsAndRefs(c CommitsAndRefs) {
	commitsAndRefs = c
}
