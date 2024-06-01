package core

import (
	"fmt"
	"gitfiend2/core/git"
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
		var repos []RepoPath
		findRepos(options.RepoPath, &repos, 0)
		return repos
	}

	return []RepoPath{}
}

const maxDepth = 5
const maxDirSize = 50

func findRepos(dir string, repos *[]RepoPath, depth int) error {
	repo, ok := getGitRepo(dir)
	if ok {
		*repos = append(*repos, repo)
		more, err := lookForSubmodules(dir)
		if err == nil {
			fmt.Println("found submodules: ", more)
		}
	}

	if depth < maxDepth {
		entries, err := os.ReadDir(dir)
		if err == nil {
			if len(entries) < maxDirSize || depth == 0 {
				for _, entry := range entries {
					if entry.IsDir() && entry.Name()[0] != '.' {
						findRepos(path.Join(dir, entry.Name()), repos, depth+1)
					}
				}
			}
		}
	}

	return nil
}

// Check if there's a submodules file and read repos from it?
func lookForSubmodules(dir string) ([]RepoPath, error) {
	file := path.Join(dir, ".gitmodules")
	_, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	text, err := os.ReadFile(file)
	var paths []RepoPath
	if err != nil {
		return nil, err
	}

	rows, ok := git.ParseConfig(string(text))
	if ok {
		for _, row := range rows {
			switch r := row.(type) {
			case git.Section:
				if r.Heading[0] == "submodule" {
					for _, sub := range r.Rows {
						switch data := sub.(type) {
						case git.DataRow:
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

type RepoPath struct {
	// TODO: What is this vs GitPath
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
