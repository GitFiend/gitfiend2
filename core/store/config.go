package store

import (
	"fmt"
	"gitfiend2/core/git"
	"os"
	"path"
)

type GitConfig struct {
	Entries    map[string]string
	Remotes    map[string]string
	Submodules map[string]string
}

func LoadConfigFromDisk(repoPath string) (GitConfig, error) {
	repo, ok := GetRepoPath(repoPath)
	if !ok {
		return GitConfig{}, fmt.Errorf("couldn't load config for %s", repoPath)
	}

	configPath := path.Join(repo.GitPath, "config")

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return GitConfig{}, err
	}

	rows, ok := git.ParseConfig(string(bytes))
	if !ok {
		return GitConfig{}, fmt.Errorf("failed to parse config %s", configPath)
	}

	remotes := map[string]string{}

	for _, row := range rows {
		switch r := row.(type) {
		case git.Section:
			if r.Heading.Key() == "remote" {
				for _, r2 := range r.Rows {
					switch r3 := r2.(type) {
					case git.DataRow:
						if r3.Key() == "url" {
							remotes[r.Heading.Value()] = r3.Value()
						}
						break
					}
				}
			}
		case git.DataRow:
		}
	}

	return GitConfig{Remotes: remotes}, nil
}
