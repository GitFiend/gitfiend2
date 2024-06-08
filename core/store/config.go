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

func (s *Store) LoadConfigFromDisk(repoPath string) (GitConfig, error) {
	repo, ok := s.GetRepoPath(repoPath)
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
	submodules := map[string]string{}
	entries := map[string]string{}

	for _, row := range rows {
		switch r := row.(type) {
		case git.Section:
			for _, e := range r.Entries() {
				entries[e[0]] = e[1]
			}

			key := r.Heading.Key()
			if key == "remote" {
				for _, r2 := range r.Rows {
					switch r3 := r2.(type) {
					case git.DataRow:
						if r3.Key() == "url" {
							remotes[r.Heading.Value()] = r3.Value()
						}
						break
					}
				}
			} else if key == "submodule" {
				for _, r2 := range r.Rows {
					switch r3 := r2.(type) {
					case git.DataRow:
						if r3.Key() == "url" {
							submodules[r.Heading.Value()] = r3.Value()
						}
						break
					}
				}
			}
		case git.DataRow:
			entries[r[0]] = r[1]
			break
		}
	}
	return GitConfig{Remotes: remotes, Submodules: submodules, Entries: entries}, nil
}

func (c *GitConfig) GetRemoteForBranch(shortName string) string {
	pushRemote, ok := c.Entries["branch."+shortName+".pushremote"]
	if ok {
		return pushRemote
	}
	pushDefault, ok := c.Entries["remote.pushdefault"]
	if ok {
		return pushDefault
	}
	remote, ok := c.Entries["branch."+shortName+".remote"]
	if ok {
		return remote
	}
	return "origin"
}

func (c *GitConfig) GetTrackingBranchName(localBranch string) string {
	remote := c.GetRemoteForBranch(localBranch)
	return "refs/remotes/" + remote + "/" + localBranch
}
