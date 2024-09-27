package git

import (
	"fmt"
	p "gitfiend2/core/parser"
	"log/slog"
	"strconv"
	"strings"
)

func commitIdsBetweenCommits(repoPath, commitId1, commitId2 string) []string {
	repo := cache.getRepo(repoPath)

	if len(repo.commits) > 0 {
		commitsMap := map[string]Commit{}
		for _, c := range repo.commits {
			commitsMap[c.Id] = c
		}
		ids, found := getCommitIdsBetweenCommitIds(commitId1, commitId2, commitsMap)
		if found {
			return ids
		}
	}

	return commitIdsBetweenCommitsFallback(repoPath, commitId1, commitId2)
}

func getCommitIdsBetweenCommitIds(
	commitId1, commitId2 string,
	commits map[string]Commit,
) (ids []string, found bool) {
	c1, have1 := commits[commitId1]
	c2, have2 := commits[commitId2]

	if !have1 || !have2 {
		return nil, false
	}

	return getCommitIdsBetweenCommits(c1, c2, commits), true
}

// Assumes a and b are in commits.
func getCommitIdsBetweenCommits(a, b Commit, commits map[string]Commit) (ids []string) {
	if a.Id == b.Id {
		return ids
	}

	aAncestors := findCommitAncestors(a, commits)
	aAncestors[a.Id] = true
	bAncestors := findCommitAncestors(b, commits)
	bAncestors[b.Id] = true

	for id := range aAncestors {
		if !bAncestors[id] {
			ids = append(ids, id)
		}
	}
	return ids
}

func countCommitsBetweenFallback(repoPath, commitId1, commitId2 string) int {
	if commitId1 == commitId2 {
		return 0
	}

	res, err := RunGit(
		RunOpts{
			RepoPath: repoPath, Args: []string{
				"rev-list",
				fmt.Sprintf("%v..%v", commitId1, commitId2),
				"--count",
			},
		},
	)

	if err == nil {
		n, err := strconv.Atoi(strings.TrimSpace(res.Stdout))
		if err == nil {
			return n
		}
	}
	return 0
}

func commitIdsBetweenCommitsFallback(repoPath, commitId1, commitId2 string) []string {
	out, err := RunGit(
		RunOpts{
			RepoPath: repoPath, Args: []string{
				"rev-list",
				fmt.Sprintf("%v..%v", commitId1, commitId2),
			},
		},
	)

	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	res, ok := p.ParseAll(PIdList, out.Stdout)
	if ok {
		return res
	}
	return nil
}

func findCommitAncestors(commit Commit, commits map[string]Commit) map[string]bool {
	ancestors := map[string]bool{}
	ancestorCommits := []Commit{commit}

	for len(ancestorCommits) > 0 {
		i := len(ancestorCommits) - 1
		c := ancestorCommits[i]
		ancestorCommits = ancestorCommits[:i]

		for _, id := range c.ParentIds {
			if !ancestors[id] {
				ancestors[id] = true
				parent, ok := commits[id]
				if ok {
					ancestorCommits = append(ancestorCommits, parent)
				}
			}
		}
	}

	return ancestors
}

type UnPushedCommits struct {
	// Commits that are un-pushed on this branch, but pushed on another.
	ThisBranch []string `json:"thisBranch"`
	// Commits that haven't been pushed period. These have more edit options available.
	AllBranches []string `json:"allBranches"`
}

func GetUnPushedCommits(repoPath string) UnPushedCommits {
	ids, found := getUnPushedCommitsComputed(repoPath)
	if found {
		all, ok := getUniqueUnPushedCommits(repoPath, ids)
		if ok {
			return UnPushedCommits{
				ThisBranch:  ids,
				AllBranches: all,
			}
		}
	} else {
		slog.Warn("getUnpushedCommits: Ref commit ids not found in commits, fallback to git request.")
	}

	res, err := RunGit(RunOpts{
		RepoPath: repoPath, Args: []string{
			"log", "HEAD", "--not", "--remotes", "--pretty=format:%H",
		},
	})
	if err != nil {
		slog.Error(err.Error())
		return UnPushedCommits{}
	}

	ids, ok := p.ParseAll(PIdList, res.Stdout)
	if ok {
		return UnPushedCommits{
			AllBranches: ids,
		}
	}

	return UnPushedCommits{}
}

func getUniqueUnPushedCommits(repoPath string, unPushedIds []string) ([]string, bool) {
	repo := cache.getRepo(repoPath)

	headRef, haveHeadRef := getHeadRef(repo.refs)
	if !haveHeadRef {
		return nil, false
	}
	remote, haveRemote := findSiblingRef(headRef, repo.refs)
	if !haveRemote {
		return nil, false
	}
	var head Commit
	for _, c := range repo.commits {
		if c.Id == headRef.CommitId {
			head = c
			break
		}
	}

	unPushedIdsUnique := map[string]bool{}
	for _, id := range unPushedIds {
		unPushedIdsUnique[id] = true
	}
	refsMap := map[string]RefInfo{}
	for _, r := range repo.refs {
		refsMap[r.Id] = r
	}
	commitsMap := map[string]Commit{}
	for _, c := range repo.commits {
		commitsMap[c.Id] = c
	}

	unique := &[]string{}
	checked := map[string]bool{}

	// TODO: Need a pointer to unique?
	unPushed(head, remote.CommitId, commitsMap, refsMap, unPushedIdsUnique, checked, unique)

	return *unique, true
}

func unPushed(
	current Commit,
	remoteId string,
	commits map[string]Commit,
	refs map[string]RefInfo,
	unPushedIds map[string]bool,
	checked map[string]bool,
	unique *[]string,
) {
	if checked[current.Id] {
		return
	}
	checked[current.Id] = true

	if current.Id == remoteId {
		for _, id := range current.Ref {
			if r, found := refs[id]; found {
				if r.RefType == Branch && r.Location == Remote {
					return
				} else if unPushedIds[current.Id] {
					*unique = append(*unique, current.Id)
				}
			}
		}
	}

	for _, id := range current.ParentIds {
		if c, found := commits[id]; found {
			unPushed(c, remoteId, commits, refs, unPushedIds, checked, unique)
		}
	}
}

func getUnPushedCommitsComputed(repoPath string) (ids []string, found bool) {
	repo := cache.getRepo(repoPath)

	commits := map[string]Commit{}
	for _, c := range repo.commits {
		commits[c.Id] = c
	}
	head, ok := getHeadRef(repo.refs)
	if !ok {
		return nil, false
	}
	remote, ok := findSiblingRef(head, repo.refs)
	if !ok {
		return nil, false
	}

	return getCommitIdsBetweenCommitIds(head.CommitId, remote.CommitId, commits)
}

func getHeadRef(refs []RefInfo) (RefInfo, bool) {
	for _, r := range refs {
		if r.Head {
			return r, true
		}
	}
	return RefInfo{}, false
}

func findSiblingRef(ref RefInfo, refs []RefInfo) (sibling RefInfo, found bool) {
	if ref.SiblingId != "" {
		for _, r := range refs {
			if ref.SiblingId == r.Id {
				return r, true
			}
		}
	}
	return RefInfo{}, false
}
