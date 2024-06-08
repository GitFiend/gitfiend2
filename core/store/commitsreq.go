package store

import (
	"cmp"
	"gitfiend2/core/git"
	"gitfiend2/core/shared"
	"slices"
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
		return s.commitsAndRefs
	}

	var commitInfo []git.CommitInfo

	if skipStashes {
		commitInfo = git.LoadCommits(repoPath, numCommits)
	} else {
		var stashes []git.CommitInfo

		reqCommits := make(chan []git.CommitInfo)
		reqStashes := make(chan []git.CommitInfo)

		go func() {
			reqCommits <- git.LoadCommits(repoPath, numCommits)
		}()
		go func() {
			reqStashes <- git.LoadStashes(repoPath)
		}()

		commitInfo, stashes = <-reqCommits, <-reqStashes
		commitInfo = append(commitInfo, stashes...)

		// TODO: Check this.
		slices.SortFunc(commitInfo, func(a, b git.CommitInfo) int {
			return cmp.Compare(a.StashId, b.StashId)
		})
	}

	commits, refs := s.convertCommitInfo(commitInfo, repoPath)
	refs = s.finishRefInfoProperties(refs, repoPath)

	result := CommitsAndRefs{
		Commits: commits,
		Refs:    refs,
	}

	s.SetCommitsAndRefs(result)
	return result
}

func (s *Store) convertCommitInfo(info []git.CommitInfo, repoPath string) ([]git.Commit, []git.RefInfo) {
	commits := make([]git.Commit, len(info))
	var refs []git.RefInfo

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

func convertCommit(info git.CommitInfo) git.Commit {
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
		Ref:        shared.Map(info.Ref, func(ref git.RefInfo) string { return ref.Id }),
		Filtered:   info.Filtered,
		NumSkipped: info.NumSkipped,
	}
}

func (s *Store) finishRefInfoProperties(refs []git.RefInfo, repoPath string) []git.RefInfo {
	c, ok := s.GetConfig(repoPath)
	if !ok {
		panic("Expected " + repoPath + " config to be already loaded")
	}

	for _, ref := range refs {
		if ref.RemoteName == "" {
			ref.RemoteName = c.GetRemoteForBranch(ref.ShortName)
		}
		ref.SiblingId = getSiblingIdForRef(ref, refs)
	}

	return refs
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
