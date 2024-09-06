package git

import "fmt"

func LoadPatches(repoPath string, commits []Commit) {
	//
}

// Normal refers to the type of commit. Commits that aren't stashes or merges.
func loadNormalPatches(repoPath string, commits []Commit, numCommits int) {
	if len(commits) > 20 {
		//
	} else {
		//
	}
}

func loadAllPatchesForNormalCommits(repoPath string, numCommits int) {
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
		return // TODO
	}

	// TODO: Parse patches
	fmt.Println(res)
}
