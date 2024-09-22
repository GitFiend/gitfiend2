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

func getCommitIdsBetweenCommitIds(commitId1, commitId2 string, commits map[string]Commit) (ids []string, found bool) {
	c1, have1 := commits[commitId1]
	c2, have2 := commits[commitId2]

	if have1 && have2 {
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

	res, err := RunGit(RunOpts{RepoPath: repoPath, Args: []string{
		"rev-list",
		fmt.Sprintf("%v..%v", commitId1, commitId2),
		"--count",
	}})

	if err == nil {
		n, err := strconv.Atoi(strings.TrimSpace(res.Stdout))
		if err == nil {
			return n
		}
	}
	return 0
}

func commitIdsBetweenCommitsFallback(repoPath, commitId1, commitId2 string) []string {
	out, err := RunGit(RunOpts{RepoPath: repoPath, Args: []string{
		"rev-list",
		fmt.Sprintf("%v..%v", commitId1, commitId2),
	}})

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
