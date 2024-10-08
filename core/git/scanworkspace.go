package git

import (
	"gitfiend2/core/shared"
	"os"
	"path"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type ScanOptions struct {
	RepoPath          string `json:"repoPath"`
	WorkspacesEnabled bool   `json:"workspacesEnabled"`
}

func ReqScanWorkspace(options ScanOptions) []string {
	res := cache.ScanWorkspace(options.RepoPath, options.WorkspacesEnabled)

	return shared.Map(
		res, func(r RepoPath) string {
			return r.Path
		},
	)
}

func (s *Cache) ScanWorkspace(repoPath string, workspacesEnabled bool) []RepoPath {
	repos := scanWorkspace(repoPath, workspacesEnabled)
	s.SetRepoPaths(repos)
	return repos
}

func scanWorkspace(repoPath string, workspacesEnabled bool) []RepoPath {
	if !workspacesEnabled {
		repo, ok := getGitRepo(repoPath)
		if ok {
			return []RepoPath{repo}
		}
	} else {
		repos := map[string]RepoPath{}
		err := findRepos(repoPath, repos, 0)
		if err == nil {
			return maps.Values(repos)
		}
	}

	return []RepoPath{}
}

const maxDepth = 5
const maxDirSize = 50

func findRepos(dir string, repos map[string]RepoPath, depth int) error {
	repo, ok := getGitRepo(dir)
	if ok {
		repos[repo.Path] = repo

		submodules, err := lookForSubmodules(dir)
		if err != nil {
			return err
		}
		if len(submodules) > 0 {
			for _, found := range submodules {
				repos[found.Path] = found
			}
		}
	}

	if depth < maxDepth {
		entries, err := os.ReadDir(dir)
		if err == nil {
			if len(entries) < maxDirSize || depth == 0 {
				for _, entry := range entries {
					if entry.IsDir() && entry.Name()[0] != '.' {
						p := path.Join(dir, entry.Name())
						if _, alreadyExists := repos[p]; !alreadyExists {
							err := findRepos(p, repos, depth+1)
							if err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}

// Check if there's a submodule file and read repos from it?
func lookForSubmodules(dir string) ([]RepoPath, error) {
	file := path.Join(dir, ".gitmodules")
	_, err := os.Stat(file)
	if err != nil {
		// No submodules to be found. Not an error.
		return nil, nil
	}

	text, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var paths []RepoPath

	rows, ok := ParseConfig(string(text))
	if ok {
		for _, row := range rows {
			switch r := row.(type) {
			case Section:
				if r.Heading[0] == "submodule" {
					for _, sub := range r.Rows {
						switch data := sub.(type) {
						case DataRow:
							if data[0] == "path" {
								repoPath, ok := getGitRepo(path.Join(dir, data[1]))
								if ok {
									paths = append(paths, repoPath)
								}
							}
						}
					}
				}
			}
		}
	}

	return paths, nil
}

// TODO: This is poorly named.
type RepoPath struct {
	Path string
	// The path to .git folder. If we have a submodule, it may be in the root repo?
	GitPath   string
	SubModule bool
}

func getGitRepo(dir string) (RepoPath, bool) {
	info, err := os.Stat(dir)
	if err != nil {
		return RepoPath{}, false
	}

	if info.IsDir() {
		gitFilePath := path.Join(dir, ".git")

		info, err := os.Stat(gitFilePath)
		if err != nil {
			return RepoPath{}, false
		}

		if !info.IsDir() {
			// If .git is a file, we have a submodule.
			p := readSubmoduleFile(gitFilePath)

			return RepoPath{
				Path:      dir,
				GitPath:   path.Join(dir, p),
				SubModule: true,
			}, true
		}
		return RepoPath{
			Path:      dir,
			GitPath:   gitFilePath,
			SubModule: false,
		}, true
	}

	return RepoPath{}, false
}

func readSubmoduleFile(filePath string) string {
	data, err := os.ReadFile(filePath)

	if err == nil {
		return parseSubmoduleFile(string(data))
	}
	return ""
}

// Expect something like:
// gitdir: ../.git/modules/cottontail-js
func parseSubmoduleFile(text string) string {
	runes := []rune(text)
	i := slices.Index(runes, ':')

	if i == -1 {
		return ""
	}
	return strings.TrimSpace(string(runes[i+1:]))
}
