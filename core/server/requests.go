package server

import (
	"encoding/json"
	"fmt"
	"gitfiend2/core/git"
	"gitfiend2/core/store"
	"os"
	"path"
)

func handleFuncRequest(name string, reqData []byte) ([]byte, bool) {
	var res any
	var ok bool

	switch name {
	case "git_version":
		res, ok = callFunc(reqGitVersion, reqData)
	case "scan_workspace":
		res, ok = callFunc(reqScanWorkspace, reqData)
	case "is_rebase_in_progress":
		res, ok = callFunc(isRebaseInProgress, reqData)
	case "load_repo_status":
		res, ok = callFunc(reqRepoStatus, reqData)
	}

	if ok {
		fmt.Println("Func Result: ", res)
		resBytes, err := json.Marshal(res)

		if err == nil {
			return resBytes, true
		}
	}

	return []byte{}, false
}

type ReqOptions struct {
	RepoPath string `json:"repoPath"`
}

type ReqResult[T any] struct {
	Ok  T
	Err error
}

func reqGitVersion(_ ReqOptions) git.VersionInfo {
	git.LoadGitVersion()
	return git.Version
}

func reqScanWorkspace(options store.ScanOptions) []string {
	res := store.ScanWorkspace(options)
	store.SetRepoPaths(res)

	var paths []string
	for _, repo := range res {
		paths = append(paths, repo.Path)
	}
	return paths
}

func reqRepoStatus(o ReqOptions) store.RepoStatus {
	return store.LoadRepoStatus(o.RepoPath)
}

func isRebaseInProgress(options ReqOptions) bool {
	p, found := store.GetRepoPath(options.RepoPath)

	if found {
		file := path.Join(p.GitPath, "rebase-merge")
		_, err := os.Stat(file)
		// Assume the file doesn't exist if we get an error.
		return err == nil
	}
	return false
}
