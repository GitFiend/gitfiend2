package store

import "gitfiend2/core/git"

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
