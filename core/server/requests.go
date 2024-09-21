package server

import (
	"encoding/json"
	"gitfiend2/core/git"
	"log/slog"
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
	case "watch_repo":
		ok = true
	case "repo_has_changed":
		ok = true
		res = false
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
		slog.Error(err.Error())
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

func loadCommitsAndRefs(o git.ReqCommitsOptions) []any {
	commitsAndRefs := git.LoadCommitsAndRefs(o)
	return []any{commitsAndRefs.Commits, commitsAndRefs.Refs}
}
