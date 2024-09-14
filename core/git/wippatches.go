package git

import (
	"fmt"
	"gitfiend2/core/git/wippatchtype"
)

type wipPatchInfo struct {
	oldFile  string
	newFile  string
	staged   wippatchtype.Type
	unStaged wippatchtype.Type
}

func LoadWipPatches(repoPath string) {
	res, err := RunGit(
		RunOpts{
			RepoPath: repoPath,
			Args: []string{
				`status`,
				`--porcelain`,
				`-uall`,
				`-z`,
			},
		},
	)

	fmt.Println(res, err)
}
