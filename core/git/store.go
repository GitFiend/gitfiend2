package git

import (
	"slices"
)

type Cache struct {
	repoPaths []RepoPath
	repos     map[string]*repo
}

type repo struct {
	config  Config
	commits []Commit
	refs    []RefInfo
	patches map[string][]Patch
}

// TODO: Use this instead of the server version.
var cache = Cache{
	repoPaths: []RepoPath{},
	repos:     map[string]*repo{},
}

func (s *Cache) SetRepoPaths(repos []RepoPath) {
	s.repoPaths = repos
}

func (s *Cache) GetRepoPath(repoPath string) (RepoPath, bool) {
	i := slices.IndexFunc(
		s.repoPaths, func(p RepoPath) bool {
			return p.Path == repoPath
		},
	)

	if i >= 0 {
		return s.repoPaths[i], true
	}
	return RepoPath{}, false
}

// TODO: May be a better way than always run this.
func (s *Cache) getRepo(repoPath string) *repo {
	r, found := s.repos[repoPath]
	if found {
		return r
	}

	r = &repo{
		config:  Config{},
		commits: make([]Commit, 0),
		refs:    make([]RefInfo, 0),
		patches: map[string][]Patch{},
	}
	s.repos[repoPath] = r

	return r
}

func (s *Cache) SetConfig(repoPath string, c Config) {
	r := s.getRepo(repoPath)
	r.config = c
}

func (s *Cache) GetConfig(repoPath string) Config {
	r := s.getRepo(repoPath)
	return r.config
}

func (s *Cache) SetCommitsAndRefs(repoPath string, c CommitsAndRefs) {
	r := s.getRepo(repoPath)
	r.commits = c.Commits
	r.refs = c.Refs
}

func (s *Cache) GetCommitsAndRefs(repoPath string) CommitsAndRefs {
	r := s.getRepo(repoPath)
	return CommitsAndRefs{Commits: r.commits, Refs: r.refs}
}
