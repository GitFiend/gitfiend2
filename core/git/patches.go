package git

import (
	"encoding/json"
	"fmt"
	"gitfiend2/core/parser"
	"os"
	"path"
	"strings"
)

func (s *Cache) LoadPatches(repoPath string, commits []Commit) map[string][]Patch {
	var (
		commitsWithoutPatches         []Commit
		stashesOrMergesWithoutPatches []Commit
		newPatches                    = map[string][]Patch{}
	)

	patches, ok := s.loadPatchesCache(repoPath)
	if ok {
		for _, c := range commits {
			if p, ok := patches[c.Id]; ok {
				newPatches[c.Id] = p
			} else if c.StashId == "" && !c.IsMerge {
				commitsWithoutPatches = append(commitsWithoutPatches, c)
			} else {
				stashesOrMergesWithoutPatches = append(stashesOrMergesWithoutPatches, c)
			}
		}
	} else {
		for _, c := range commits {
			if c.StashId == "" && !c.IsMerge {
				commitsWithoutPatches = append(commitsWithoutPatches, c)
			} else {
				stashesOrMergesWithoutPatches = append(stashesOrMergesWithoutPatches, c)
			}
		}
	}

	if len(commitsWithoutPatches) == 0 && len(stashesOrMergesWithoutPatches) == 0 {
		return patches
	}

	if len(commitsWithoutPatches) > 0 {
		p, ok := loadNormalPatches(repoPath, commitsWithoutPatches, len(commits))
		if ok {
			for id, patch := range p {
				patches[id] = patch
			}
		}
	}

	for _, c := range stashesOrMergesWithoutPatches {
		p := loadPatchesForCommit(repoPath, c)
		newPatches[c.Id] = p
	}
	writePatchesCache(repoPath, newPatches)
	return newPatches
}

func writePatchesCache(repoPath string, newPatches map[string][]Patch) {
	file, ok := getCacheFile(repoPath)
	if !ok {
		return
	}

	data, err := json.Marshal(newPatches)

	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func loadPatchesForCommit(repoPath string, commit Commit) []Patch {
	var args []string

	if commit.IsMerge {
		args = []string{
			"diff",
			"--name-status",
			"-z",
			"--no-color",
			fmt.Sprintf("%s...%s", commit.ParentIds[0], commit.ParentIds[1]),
		}
	} else if commit.StashId != "" {
		args = []string{
			"diff",
			fmt.Sprintf("%s..%s", commit.ParentIds[0], commit.Id),
			"--no-color",
			"--name-status",
			"-z",
		}
	} else {
		args = []string{
			"diff", fmt.Sprintf("%s..%s", commit0Id, commit.Id),
			"--no-color",
			"--name-status",
			"-z",
		}
	}

	res, err := RunGit(
		RunOpts{
			RepoPath: repoPath,
			Args:     args,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	patches, ok := parser.ParseAll(pPatches, res.Stdout)
	if ok {
		return patchDataToPatches(commit.Id, patches)
	}
	return nil
}

func (s *Cache) loadPatchesCache(repoPath string) (map[string][]Patch, bool) {
	r := s.ensureRepo(repoPath)
	if len(r.patches) > 0 {
		return r.patches, true
	}

	file, ok := getCacheFile(repoPath)
	if !ok {
		return nil, false
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, false
	}

	var patches map[string][]Patch
	err = json.Unmarshal(bytes, &patches)
	r.patches = patches

	return patches, err == nil
}

func getCacheFile(repoPath string) (string, bool) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", false
	}

	appDir := path.Join(dir, "gitfiend")
	err = os.MkdirAll(appDir, os.ModeDir)
	if err != nil {
		return "", false
	}

	return path.Join(appDir, genCacheFileName(repoPath)), true
}

func genCacheFileName(repoPath string) string {
	r := strings.NewReplacer("\\", "", ":", "", "/", "")
	fileName := r.Replace(repoPath)

	return fileName + ".json"
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
