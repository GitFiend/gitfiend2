package server

import (
	"encoding/json"
	"gitfiend2/core/git"
	"github.com/labstack/gommon/log"
)

func handleFuncRequest(name string, reqData []byte) ([]byte, bool) {
	var res any
	var ok bool

	switch name {
	case "git_version":
		res, ok = callFunc(reqGitVersion, reqData)
	case "scan_workspace":
		res, ok = callFunc(git.ReqScanWorkspace, reqData)
	case "is_rebase_in_progress":
		res, ok = callFunc(isRebaseInProgress, reqData)
	case "load_repo_status":
		res, ok = callFunc(reqRepoStatus, reqData)
	case "load_commits_and_refs":
		res, ok = callFunc(loadCommitsAndRefs, reqData)
	case "load_wip_patches":
		res, ok = callFunc(loadWipPatches, reqData)
	}

	if ok {
		//fmt.Println("Func Result: ", res)
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

func loadWipPatches(o ReqOptions) git.WipPatches {
	res, err := git.LoadWipPatches(o.RepoPath)
	if err != nil {
		log.Error(err)
	}
	return res
}

func reqGitVersion(_ ReqOptions) git.VersionInfo {
	git.LoadGitVersion()
	return git.Version
}

func reqRepoStatus(o ReqOptions) git.RepoStatus {
	return git.LoadRepoStatus(o.RepoPath)
}

func isRebaseInProgress(options ReqOptions) bool {
	return git.IsRebaseInProgress(options.RepoPath)
}

func loadCommitsAndRefs(o git.ReqCommitsOptions) git.CommitsAndRefs {
	return git.LoadCommitsAndRefs(o)
}
