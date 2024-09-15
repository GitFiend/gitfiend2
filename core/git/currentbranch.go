package git

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// TODO: Needs testing.
func loadCurrentBranch(repoPath string) (id string, name string, ok bool) {
	rp, ok := cache.GetRepoPath(repoPath)
	if !ok {
		return
	}

	head := path.Join(rp.GitPath, "HEAD")

	data, err := os.ReadFile(head)
	if err != nil {
		fmt.Println(err)
	}

	text := string(data)
	i := strings.LastIndex(text, ":")
	if i > 0 {
		id = strings.TrimSpace(text[i:])
		name = strings.Replace(id, "refs/heads/", "", 1)
		ok = true
		return
	}
	return
}

func readRefs(repoPath, branchName string) {
	rp, ok := cache.GetRepoPath(repoPath)
	if !ok {
		panic("Can't read refs without a repo path")
	}

	refsDir := path.Join(rp.GitPath, "refs")
	headsDir := path.Join(refsDir, "heads")
}
