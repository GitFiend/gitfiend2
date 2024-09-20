package git

import (
	"os"
	"path"
)

type RepoStatus struct {
	Patches        WipPatches `json:"patches"`
	Config         Config     `json:"config"`
	Branches       []string   `json:"branches"`
	RemoteAhead    int        `json:"remoteAhead"`
	RemoteBehind   int        `json:"remoteBehind"`
	BranchName     string     `json:"branchName"`
	HeadRefId      string     `json:"headRefId"`
	LocalCommitId  string     `json:"localCommitId"`
	RemoteCommitId string     `json:"remoteCommitId"`
	State          string     `json:"state"`
}

func LoadRepoStatus(repoPath string) RepoStatus {
	patches, err := LoadWipPatches(repoPath)
	if err != nil {
		panic("Failed to load patches")
	}

	config := cache.LoadFullConfig(repoPath)
	headId, currentBranch, ok := loadCurrentBranch(repoPath)
	if !ok {
		panic("Failed to load current branch")
	}

	refs := readRefs(repoPath, currentBranch)

	packedRefs, ok := loadPackedRefs(repoPath)
	if !ok {
		// TODO: Is this actually ok?
		panic("Failed to load packed refs")
	}

	if refs.localId == "" {
		for _, r := range packedRefs {
			if r.RemoteName == "" && r.Name == currentBranch {
				refs.localId = r.CommitId
				break
			}
		}
	}
	if refs.remoteId == "" {
		for _, r := range packedRefs {
			if r.Name == currentBranch {
				refs.remoteId = r.CommitId
				break
			}
		}
	}
	for _, r := range packedRefs {
		if r.Name != "" {
			refs.others[r.Name] = true
		}
	}

	if refs.localId != "" {
		if refs.remoteId != "" {
			remoteAhead := countCommitsBetweenFallback(repoPath, refs.localId, refs.remoteId)
			remoteBehind := countCommitsBetweenFallback(repoPath, refs.remoteId, refs.localId)
			var branches []string
			for name := range refs.others {
				branches = append(branches, name)
			}

			return RepoStatus{
				Patches:        patches,
				Config:         config,
				RemoteAhead:    remoteAhead,
				RemoteBehind:   remoteBehind,
				Branches:       branches,
				BranchName:     currentBranch,
				HeadRefId:      headId,
				LocalCommitId:  refs.localId,
				RemoteCommitId: refs.remoteId,
				State:          "Both",
			}
		}
	}

	var branches []string
	for name := range refs.others {
		branches = append(branches, name)
	}

	// TODO: Check this.
	state := "Local"
	if refs.localId == "" {
		state = "Remote"
	} else if refs.remoteId != "" {
		state = "Both"
	} else if refs.others["HEAD"] {
		state = "Remote"
	}

	return RepoStatus{
		Patches:        patches,
		Config:         config,
		RemoteAhead:    0,
		RemoteBehind:   0,
		Branches:       branches,
		BranchName:     currentBranch,
		HeadRefId:      headId,
		LocalCommitId:  refs.localId,
		RemoteCommitId: refs.remoteId,
		State:          state,
	}
}

func IsRebaseInProgress(repoPath string) bool {
	p, found := cache.GetRepoPath(repoPath)

	if found {
		file := path.Join(p.GitPath, "rebase-merge")
		_, err := os.Stat(file)
		// Assume the file doesn't exist if we get an error.
		return err == nil
	}
	return false
}
