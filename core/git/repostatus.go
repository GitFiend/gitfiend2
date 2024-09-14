package git

import "fmt"

type RepoStatus struct {
	Config `json:"config"`
}

func (s *Store) LoadRepoStatus(repoPath string) RepoStatus {
	config := s.LoadFullConfig(repoPath)

	id, name, ok := s.loadCurrentBranch(repoPath)

	fmt.Println(id, name, ok)

	return RepoStatus{Config: config}
}
