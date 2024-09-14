package git

import (
	"slices"
)

type Store struct {
	repoPaths []RepoPath
	repos     map[string]*repo
}

type repo struct {
	config  Config
	commits []Commit
	refs    []RefInfo
	patches map[string][]Patch
}

func NewStore() Store {
	return Store{
		repoPaths: []RepoPath{},
		repos:     map[string]*repo{},
	}
}

// TODO: Use this instead of the server version.
var store = NewStore()

func (s *Store) SetRepoPaths(repos []RepoPath) {
	s.repoPaths = repos
}

func (s *Store) GetRepoPath(repoPath string) (RepoPath, bool) {
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
func (s *Store) ensureRepo(repoPath string) *repo {
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

func (s *Store) SetConfig(repoPath string, c Config) {
	r := s.ensureRepo(repoPath)
	r.config = c
}

func (s *Store) GetConfig(repoPath string) Config {
	r := s.ensureRepo(repoPath)
	return r.config
}

func (s *Store) SetCommitsAndRefs(repoPath string, c CommitsAndRefs) {
	r := s.ensureRepo(repoPath)
	r.commits = c.Commits
	r.refs = c.Refs
}

func (s *Store) GetCommitsAndRefs(repoPath string) CommitsAndRefs {
	r := s.ensureRepo(repoPath)
	return CommitsAndRefs{Commits: r.commits, Refs: r.refs}
}
