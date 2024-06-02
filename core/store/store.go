package store

import (
	"gitfiend2/core"
	"slices"
)

var RepoPaths = []core.RepoPath{}

func SetRepoPaths(repos []core.RepoPath) {
	RepoPaths = repos
}

func GetRepoPath(repoPath string) (core.RepoPath, bool) {
	i := slices.IndexFunc(RepoPaths, func(p core.RepoPath) bool {
		return p.Path == repoPath
	})

	if i >= 0 {
		return RepoPaths[i], true
	}
	return core.RepoPath{}, false
}
