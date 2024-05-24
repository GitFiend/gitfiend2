package core

import (
	"os"
	"path"
	"slices"
	"strings"
)

type ScanOptions struct {
	RepoPath          string
	WorkspacesEnabled bool
}

func ScanWorkspace(options ScanOptions) []RepoPath {
	if !options.WorkspacesEnabled {
		repo, ok := getGitRepo(options.RepoPath)
		if ok {
			return []RepoPath{repo}
		}
	} else {
		//
	}

	return []RepoPath{}
}

const maxDepth = 5
const maxDirSize = 50

func findRepos(dir string, repos []RepoPath, depth int) {
	repo, ok := getGitRepo(dir)
	if ok {
		repos = append(repos, repo)
	}

	if depth < maxDepth {
		// TODO: Check if this dir contains .gitmodules file.

		entries, err := os.ReadDir(dir)
		if err != nil {
			if len(entries) < maxDirSize || depth == 0 {
				for _, entry := range entries {
					if entry.IsDir() && entry.Name()[0] != '.' {
						findRepos(path.Join(dir, entry.Name()), repos, depth+1)
					}
				}
			}
		}
	}
}

type RepoPath struct {
	Path      string
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
				Path:      p,
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

	if err != nil {
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
