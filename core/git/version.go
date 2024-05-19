package git

import (
	"strconv"
	"strings"
	"unicode"
)

type VersionInfo struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

func (v VersionInfo) Valid() bool {
	return v.Major > 0
}

var Version VersionInfo

func LoadGitVersion() {
	res, err := RunGit(RunOpts{
		RepoPath: ".",
		Args:     []string{"--version"},
	})

	if err != nil {
		return
	}

	Version = ParseGitVersion(res.Stdout)
}

func ParseGitVersion(s string) VersionInfo {
	i := strings.IndexFunc(s, func(r rune) bool {
		return unicode.IsDigit(r)
	})

	j := i
	for _, c := range s[i:] {
		if !unicode.IsDigit(c) && c != '.' {
			break
		}
		j++
	}

	sub := s[i:j]
	parts := strings.Split(sub, ".")

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])

	var patch int
	if len(parts) > 2 {
		patch, _ = strconv.Atoi(parts[2])
	}

	return VersionInfo{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}
