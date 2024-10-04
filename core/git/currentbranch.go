package git

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"slices"
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
		id = strings.TrimSpace(text[i+1:])
		name = strings.Replace(id, "refs/heads/", "", 1)
		ok = true
		return
	}
	return
}

func readRefs(repoPath, branchName string) *refs {
	rp, ok := cache.GetRepoPath(repoPath)
	if !ok {
		panic("Can't read refs without a repo path")
	}

	refsDir := path.Join(rp.GitPath, "refs")
	headsDir := path.Join(refsDir, "heads")

	refs := &refs{
		others: map[string]bool{},
	}
	readLocalRefs(headsDir, headsDir, branchName, refs)

	remotesDir := path.Join(refsDir, "remotes")
	remotes, err := os.ReadDir(remotesDir)
	// Sometimes remotes folder doesn't exist.
	if err == nil {
		for _, e := range remotes {
			p := path.Join(remotesDir, e.Name())
			readRemoteRefs(p, p, branchName, refs)
		}
	}
	return refs
}

type refs struct {
	localId, remoteId string
	others            map[string]bool
}

func readLocalRefs(currentDir, startDir, branchName string, refs *refs) {
	entities, err := os.ReadDir(currentDir)
	if err != nil {
		panic("Can't read local refs")
	}

	for _, e := range entities {
		name := e.Name()
		p := path.Join(currentDir, name)
		if e.IsDir() {
			readLocalRefs(p, startDir, branchName, refs)
		} else if name[0] != '.' && name != "HEAD" {
			foundRef := getRefNameFromPath(p, startDir)
			if foundRef == branchName {
				refs.localId = readIdFromRefPath(p)
			} else {
				refs.others[foundRef] = true
			}
		}
	}
}

func readRemoteRefs(currentDir, startDir, branchName string, refs *refs) {
	entities, err := os.ReadDir(currentDir)
	if err != nil {
		slog.Warn("Remote refs dir not found (" + currentDir + ")")
		return
	}

	for _, e := range entities {
		name := e.Name()
		p := path.Join(currentDir, name)
		if e.IsDir() {
			readRemoteRefs(p, startDir, branchName, refs)
		} else if name[0] != '.' {
			if name == "HEAD" {
				if strings.HasSuffix(readHeadFile(p), branchName) {
					refs.remoteId = refs.localId
				}
			} else {
				foundRef := getRefNameFromPath(p, startDir)
				if foundRef == branchName {
					refs.remoteId = readIdFromRefPath(p)
				} else {
					refs.others[foundRef] = true
				}
			}
		}
	}
}

func getRefNameFromPath(filePath, repoDir string) string {
	filePath = strings.Replace(filePath, repoDir, "", 1)
	dir, file := path.Split(filePath)
	parts := []string{file}

	for dir != "" {
		dir, file = path.Split(dir)
		if len(dir) > 0 {
			dir = dir[:len(dir)-1]
		}
		if len(file) > 0 {
			parts = append(parts, file)
		}
	}
	slices.Reverse(parts)
	return strings.Join(parts, "/")
}

func readIdFromRefPath(refPath string) string {
	bytes, err := os.ReadFile(refPath)
	if err != nil {
		slog.Error("Failed to read " + refPath + " for id")
	}
	text := string(bytes)
	return strings.TrimSpace(text)
}

func readHeadFile(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Failed to read" + filePath)
	}

	text := string(bytes)
	i := strings.IndexRune(text, ':')
	if i >= 0 {
		p := text[i+1:]
		return strings.TrimSpace(p)
	}

	slog.Error("Didn't find ':' in the head file")
	return ""
}
