package git

import (
	"fmt"
	"gitfiend2/core/parser"
)

func LoadPatches(repoPath string, commits []Commit) {
	//
}

// Normal refers to the type of commit. Commits that aren't stashes or merges.
func loadNormalPatches(repoPath string, commits []Commit, numCommits int) (
	map[string][]Patch,
	bool,
) {
	if len(commits) > 20 {
		return loadAllPatchesForNormalCommits(repoPath, numCommits)
	} else {
		args := []string{"show"}
		for _, c := range commits {
			args = append(args, c.Id)
		}
		args = append(args, "--name-status", "--pretty=format:%H", "-z")

		res, err := RunGit(
			RunOpts{
				RepoPath: repoPath,
				Args:     args,
			},
		)
		if err != nil {
			return nil, false
		}

		return parser.ParseAll(pManyPatchesWithCommitIds, res.Stdout)
	}
}

func loadAllPatchesForNormalCommits(repoPath string, numCommits int) (
	map[string][]Patch,
	bool,
) {
	res, err := RunGit(
		RunOpts{
			Args: []string{
				"log",
				"--remotes",
				"--branches",
				"--all",
				"--name-status",
				"--pretty=format:%H,",
				// Can't get the correct patch info for merges with this command.
				"--no-merges",
				"-z",
				fmt.Sprintf("-n%d", numCommits),
			},
			RepoPath: repoPath,
		},
	)

	if err != nil {
		return nil, false
	}

	return parser.ParseAll(pManyPatchesWithCommitIds, res.Stdout)
}
