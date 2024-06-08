package store

type RepoStatus struct {
}

func LoadRepoStatus(repoPath string) RepoStatus {
	c, err := LoadConfigFromDisk(repoPath)

	if err == nil {
		SetConfig(repoPath, c)
	}

	return RepoStatus{}
}
