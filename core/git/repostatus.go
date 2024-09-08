package git

type RepoStatus struct {
	Config Config `json:"config"`
}

func (s *Store) LoadRepoStatus(repoPath string) RepoStatus {
	config := s.LoadFullConfig(repoPath)

	return RepoStatus{Config: config}
}
