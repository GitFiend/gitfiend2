package git

import (
	"fmt"
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

func parseGitVersion(s string) VersionInfo {
	for _, r := range s {
		if unicode.IsDigit(r) {
			//
		}
	}

	return VersionInfo{}
}
