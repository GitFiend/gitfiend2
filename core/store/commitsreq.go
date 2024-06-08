package store

import (
	"gitfiend2/core/git"
	"gitfiend2/core/shared"
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
	c, ok := GetConfig(repoPath)
	if !ok {
		panic("Expected " + repoPath + " config to be already loaded")
	}

	for _, ref := range refs {
		if ref.RemoteName == "" {
			ref.RemoteName = c.GetRemoteForBranch(ref.ShortName)
		}
		ref.SiblingId = getSiblingIdForRef(ref, refs)
	}
}

func getSiblingIdForRef(ref git.RefInfo, refs []git.RefInfo) string {
	if ref.Location == git.Remote {
		local, ok := shared.Find(refs, func(r git.RefInfo) bool {
			return r.Location == git.Local && r.ShortName == ref.ShortName
		})
		if ok {
			return local.Id
		}
	}
	remote, ok := shared.Find(refs, func(r git.RefInfo) bool {
		return r.Location == git.Remote &&
			r.ShortName == ref.ShortName &&
			r.RemoteName == ref.RemoteName
	})
	if ok {
		return remote.Id
	}
	return ""
}
