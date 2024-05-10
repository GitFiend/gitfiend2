package git

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type VersionInfo struct {
	Major int
	Minor int
	Patch int
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

	// TODO
	fmt.Println(res)
}

func ParseGitVersion(s string) VersionInfo {
	i := strings.IndexFunc(s, func(r rune) bool {
		return unicode.IsDigit(r)
	})
	j := strings.LastIndexFunc(s[i:], func(r rune) bool {
		return unicode.IsDigit(r)
	})

	sub := s[i : i+j+1]

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
