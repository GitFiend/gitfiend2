package core

import (
	"os"
	"path"
)

type ScanOptions struct {
	RepoPath          string
	WorkspacesEnabled bool
}

func ScanWorkspace(options ScanOptions) {
	//
}

func scanSingleRepo(dir string) {
	//
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
			//
		}
		return RepoPath{
			Path:      dir,
			GitPath:   gitFilePath,
			SubModule: false,
		}, true
	}

	return RepoPath{}, false
}
