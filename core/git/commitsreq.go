package git

import (
	"cmp"
	"gitfiend2/core/shared"
	"slices"
	"strings"
)

type ReqCommitsOptions struct {
	RepoPath    string         `json:"repoPath"`
	NumCommits  int            `json:"numCommits"`
	Filters     []CommitFilter `json:"filters"`
	Fast        bool           `json:"fast"`
	SkipStashes bool           `json:"skipStashes"`
}

type CommitsAndRefs struct {
	Commits []Commit  `json:"commits"`
	Refs    []RefInfo `json:"refs"`
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
	skipStashes bool,
) CommitsAndRefs {
	if cacheOnly {
		return cache.GetCommitsAndRefs(repoPath)
	}

	var commitInfo []CommitInfo

	if skipStashes {
		commitInfo = LoadCommits(repoPath, numCommits)
	} else {
		var stashes []CommitInfo

		reqCommits := make(chan []CommitInfo)
		reqStashes := make(chan []CommitInfo)

		go func() {
			reqCommits <- LoadCommits(repoPath, numCommits)
		}()
		go func() {
			reqStashes <- LoadStashes(repoPath)
		}()

		commitInfo, stashes = <-reqCommits, <-reqStashes
		commitInfo = append(commitInfo, stashes...)

		slices.SortFunc(
			commitInfo, func(a, b CommitInfo) int {
				if a.StashId != "" || b.StashId != "" {
					return cmp.Compare(b.Date.Ms, a.Date.Ms)
				}
				return 0
			},
		)
	}

	commits, refs := convertCommitInfo(commitInfo, repoPath)

	result := CommitsAndRefs{
		Commits: commits,
		Refs:    refs,
	}

	// TODO: Set indices on commits.

	cache.SetCommitsAndRefs(repoPath, result)
	return result
}

func convertCommitInfo(info []CommitInfo, repoPath string) (
	[]Commit,
	[]RefInfo,
) {
	commits := make([]Commit, len(info))
	var refs []RefInfo

	for i, c := range info {
		c.Index = i
		for _, r := range c.Ref {
			if !strings.Contains(r.FullName, "HEAD") {
				refs = append(refs, r)
			}
		}
		commits[i] = convertCommit(c)
	}
	return commits, finishRefInfoProperties(refs, repoPath)
}

func convertCommit(info CommitInfo) Commit {
	return Commit{
		Author:     info.Author,
		Email:      info.Email,
		Date:       info.Date,
		Id:         info.Id,
		Index:      info.Index,
		ParentIds:  info.ParentIds,
		IsMerge:    info.IsMerge,
		Message:    info.Message,
		StashId:    info.StashId,
		Ref:        shared.Map(info.Ref, func(ref RefInfo) string { return ref.Id }),
		Filtered:   info.Filtered,
		NumSkipped: info.NumSkipped,
	}
}

func finishRefInfoProperties(refs []RefInfo, repoPath string) []RefInfo {
	c := cache.GetConfig(repoPath)

	for i := range refs {
		ref := &refs[i]
		if ref.RemoteName == "" {
			ref.RemoteName = c.GetRemoteForBranch(ref.ShortName)
		}
		ref.SiblingId = getSiblingIdForRef(ref, refs)
	}

	return refs
}

func getSiblingIdForRef(ref *RefInfo, refs []RefInfo) string {
	if ref.Location == Remote {
		local, ok := shared.Find(
			refs, func(r RefInfo) bool {
				return r.Location == Local && r.ShortName == ref.ShortName
			},
		)
		if ok {
			return local.Id
		}
		return ""
	}
	remote, ok := shared.Find(
		refs, func(r RefInfo) bool {
			return r.Location == Remote &&
				r.ShortName == ref.ShortName &&
				r.RemoteName == ref.RemoteName
		},
	)
	if ok {
		return remote.Id
	}
	return ""
}
