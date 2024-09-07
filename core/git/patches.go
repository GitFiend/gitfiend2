package git

import (
	"encoding/json"
	"fmt"
	"gitfiend2/core/parser"
	"os"
	"path"
	"strings"
)

func (s *Store) LoadPatches(repoPath string, commits []Commit) map[string][]Patch {
	patches, ok := s.loadPatchesCache(repoPath)
	if ok {
		return patches
	}

	// TODO
	return nil
}

func (s *Store) loadPatchesCache(repoPath string) (map[string][]Patch, bool) {
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
