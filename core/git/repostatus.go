package git

import (
	"fmt"
	"os"
	"path"
)

type RepoStatus struct {
	Config `json:"config"`
}

func LoadRepoStatus(repoPath string) RepoStatus {
	patches, err := LoadWipPatches(repoPath)
	if err != nil {
		panic("Failed to load patches")
	}
	fmt.Println(patches)

	config := cache.LoadFullConfig(repoPath)
	id, name, ok := loadCurrentBranch(repoPath)

	fmt.Println(id, name, ok)

	refs := readRefs(repoPath, name)
	fmt.Println(refs)

	return RepoStatus{Config: config}
}

func IsRebaseInProgress(repoPath string) bool {
	p, found := cache.GetRepoPath(repoPath)

	if found {
		file := path.Join(p.GitPath, "rebase-merge")
		_, err := os.Stat(file)
		// Assume the file doesn't exist if we get an error.
		return err == nil
	}
	return false
}
