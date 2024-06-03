package server

import (
	"gitfiend2/core"
	"gitfiend2/core/git"
	"gitfiend2/core/store"
	"os"
	"path"
)

type ReqOptions struct {
	RepoPath string `json:"repoPath"`
}

type ReqResult[T any] struct {
	Ok  T
	Err error
}

func ReqGitVersion(_ ReqOptions) git.VersionInfo {
	git.LoadGitVersion()
	return git.Version
}

func ReqScanWorkspace(options core.ScanOptions) []string {
	res := core.ScanWorkspace(options)

	store.SetRepoPaths(res)

	var paths []string
	for _, repo := range res {
		paths = append(paths, repo.Path)
	}
	return paths
}

func IsRebaseInProgress(options ReqOptions) bool {
	p, found := store.GetRepoPath(options.RepoPath)

	if found {
		file := path.Join(p.GitPath, "rebase-merge")
		_, err := os.Stat(file)
		// Assume the file doesn't exist if we get an error.
		return err == nil
	}
	return false
}
