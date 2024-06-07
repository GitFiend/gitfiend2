package store

import (
	"gitfiend2/core/git"
	"strings"
)

type ReqCommitsOptions struct {
	RepoPath    string             `json:"repoPath"`
	NumCommits  int                `json:"numCommits"`
	Filters     []git.CommitFilter `json:"filters"`
	Fast        bool               `json:"fast"`
	SkipStashes bool               `json:"skipStashes"`
}

type CommitsAndRefs struct {
	Commits []git.Commit  `json:"commits"`
	Refs    []git.RefInfo `json:"refs"`
}

func LoadCommitsAndRefs(o ReqCommitsOptions) CommitsAndRefs {
	res := loadCommitsUnfiltered(o.RepoPath, o.NumCommits, o.Fast, o.SkipStashes)

	// TODO: Filters.
	return res
}

func loadCommitsUnfiltered(
	repoPath string,
	numCommits int,
	cacheOnly bool,
	skipStashes bool) CommitsAndRefs {
	if cacheOnly {
		return commitsAndRefs
	}

	if skipStashes {
		// load commits
	}

	return commitsAndRefs
}

func convertCommitInfo(info []git.CommitInfo) ([]git.Commit, []git.RefInfo) {
	commits := make([]git.Commit, len(info))
	var refs []git.RefInfo

	for i, c := range info {
		commits[i] = convertCommit(c)
		for _, r := range c.Ref {
			if !strings.Contains(r.FullName, "HEAD") {
				refs = append(refs, r)
			}
		}
	}
	return commits, refs
}

func convertCommit(info git.CommitInfo) git.Commit {
	refIds := make([]string, len(info.Ref))
	for i, ref := range info.Ref {
		refIds[i] = ref.Id
	}
	return git.Commit{
		Author:     info.Author,
		Email:      info.Email,
		Date:       info.Date,
		Id:         info.Id,
		Index:      info.Index,
		ParentIds:  info.ParentIds,
		IsMerge:    info.IsMerge,
		Message:    info.Message,
		StashId:    info.StashId,
		Ref:        refIds,
		Filtered:   info.Filtered,
		NumSkipped: info.NumSkipped,
	}
}

func finishRefInfoProperties(refs []git.RefInfo, repoPath string) {
	//
}
