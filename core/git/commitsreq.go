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

func (s *Store) LoadCommitsAndRefs(o ReqCommitsOptions) CommitsAndRefs {
	res := s.loadCommitsUnfiltered(o.RepoPath, o.NumCommits, o.Fast, o.SkipStashes)

	// TODO: Filters.
	return res
}

func (s *Store) loadCommitsUnfiltered(
	repoPath string,
	numCommits int,
	cacheOnly bool,
	skipStashes bool,
) CommitsAndRefs {
	if cacheOnly {
		return s.GetCommitsAndRefs(repoPath)
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

		// TODO: Check this.
		slices.SortFunc(
			commitInfo, func(a, b CommitInfo) int {
				return cmp.Compare(a.StashId, b.StashId)
			},
		)
	}

	commits, refs := s.convertCommitInfo(commitInfo, repoPath)
	refs = s.finishRefInfoProperties(refs, repoPath)

	result := CommitsAndRefs{
		Commits: commits,
		Refs:    refs,
	}

	s.SetCommitsAndRefs(repoPath, result)
	return result
}

func (s *Store) convertCommitInfo(info []CommitInfo, repoPath string) (
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
	return commits, s.finishRefInfoProperties(refs, repoPath)
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

func (s *Store) finishRefInfoProperties(refs []RefInfo, repoPath string) []RefInfo {
	c := s.GetConfig(repoPath)

	for _, ref := range refs {
		if ref.RemoteName == "" {
			ref.RemoteName = c.GetRemoteForBranch(ref.ShortName)
		}
		ref.SiblingId = getSiblingIdForRef(ref, refs)
	}

	return refs
}

func getSiblingIdForRef(ref RefInfo, refs []RefInfo) string {
	if ref.Location == Remote {
		local, ok := shared.Find(
			refs, func(r RefInfo) bool {
				return r.Location == Local && r.ShortName == ref.ShortName
			},
		)
		if ok {
			return local.Id
		}
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
