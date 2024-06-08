package store

import (
	"slices"
)

type Store struct {
	repoPaths      []RepoPath
	commitsAndRefs CommitsAndRefs
	configs        map[string]GitConfig
}

func New() Store {
	return Store{
		repoPaths:      []RepoPath{},
		commitsAndRefs: CommitsAndRefs{},
		configs:        map[string]GitConfig{},
	}
}

func (s *Store) SetRepoPaths(repos []RepoPath) {
	s.repoPaths = repos
}

func (s *Store) GetRepoPath(repoPath string) (RepoPath, bool) {
	i := slices.IndexFunc(s.repoPaths, func(p RepoPath) bool {
		return p.Path == repoPath
	})

	if i >= 0 {
		return s.repoPaths[i], true
	}
	return RepoPath{}, false
}

func (s *Store) SetConfig(repoPath string, c GitConfig) {
	s.configs[repoPath] = c
}

func (s *Store) GetConfig(repoPath string) (GitConfig, bool) {
	c, ok := s.configs[repoPath]
	return c, ok
}

func (s *Store) SetCommitsAndRefs(c CommitsAndRefs) {
	s.commitsAndRefs = c
}
