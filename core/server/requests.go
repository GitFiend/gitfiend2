package server

import (
	"gitfiend2/core"
	"gitfiend2/core/git"
)

type ReqOptions struct {
	RepoPath string `json:"repoPath"`
}

func ReqGitVersion(_ ReqOptions) git.VersionInfo {
	git.LoadGitVersion()
	return git.Version
}

func ReqScanWorkspace(options core.ScanOptions) []string {
	res := core.ScanWorkspace(options)

	var paths []string
	for _, repo := range res {
		paths = append(paths, repo.Path)
	}
	return paths
}
