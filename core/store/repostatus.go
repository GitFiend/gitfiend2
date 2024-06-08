package store

type RepoStatus struct {
}

func (s *Store) LoadRepoStatus(repoPath string) RepoStatus {
	c, err := s.LoadConfigFromDisk(repoPath)

	if err == nil {
		s.SetConfig(repoPath, c)
	}

	return RepoStatus{}
}
