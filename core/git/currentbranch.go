package git

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// TODO: Needs testing.
func (s *Cache) loadCurrentBranch(repoPath string) (id string, name string, ok bool) {
	rp, ok := s.GetRepoPath(repoPath)
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
