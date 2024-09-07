package git

import (
	"fmt"
	"os"
	"path"
)

type Config struct {
	Entries    map[string]string
	Remotes    map[string]string
	Submodules map[string]string
}

func (s *Store) LoadConfigFromDisk(repoPath string) (Config, error) {
	repo, ok := s.GetRepoPath(repoPath)
	if !ok {
		return Config{}, fmt.Errorf("couldn't load config for %s", repoPath)
	}

	configPath := path.Join(repo.GitPath, "config")

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	rows, ok := ParseConfig(string(bytes))
	if !ok {
		return Config{}, fmt.Errorf("failed to parse config %s", configPath)
	}

	remotes := map[string]string{}
	submodules := map[string]string{}
	entries := map[string]string{}

	for _, row := range rows {
		switch r := row.(type) {
		case Section:
			for _, e := range r.Entries() {
				entries[e[0]] = e[1]
			}

			key := r.Heading.Key()
			if key == "remote" {
				for _, r2 := range r.Rows {
					switch r3 := r2.(type) {
					case DataRow:
						if r3.Key() == "url" {
							remotes[r.Heading.Value()] = r3.Value()
						}
						break
					}
				}
			} else if key == "submodule" {
				for _, r2 := range r.Rows {
					switch r3 := r2.(type) {
					case DataRow:
						if r3.Key() == "url" {
							submodules[r.Heading.Value()] = r3.Value()
						}
						break
					}
				}
			}
		case DataRow:
			entries[r[0]] = r[1]
			break
		}
	}
	return Config{Remotes: remotes, Submodules: submodules, Entries: entries}, nil
}

func (c *Config) GetRemoteForBranch(shortName string) string {
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

func (c *Config) GetTrackingBranchName(localBranch string) string {
	remote := c.GetRemoteForBranch(localBranch)
	return "refs/remotes/" + remote + "/" + localBranch
}
