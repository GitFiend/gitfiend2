package store

import (
	"slices"
)

var repoPaths []RepoPath
var commitsAndRefs CommitsAndRefs
var configs = map[string]GitConfig{}

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

func SetConfig(repoPath string, c GitConfig) {
	configs[repoPath] = c
}
func GetConfig(repoPath string) (GitConfig, bool) {
	c, ok := configs[repoPath]
	return c, ok
}

func SetCommitsAndRefs(c CommitsAndRefs) {
	commitsAndRefs = c
}
