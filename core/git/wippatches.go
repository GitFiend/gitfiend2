package git

import (
	"cmp"
	"fmt"
	"gitfiend2/core/git/wippatchtype"
	p "gitfiend2/core/parser"
	"github.com/labstack/gommon/log"
	"os"
	"path"
	"slices"
	"strings"
)

type wipPatchInfo struct {
	oldFile  string
	newFile  string
	staged   wippatchtype.Type
	unStaged wippatchtype.Type
}

func LoadWipPatches(repoPath string) (WipPatches, error) {
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
	if err != nil {
		return WipPatches{}, err
	}

	info, ok := p.ParseAll(pWipPatches, res.Stdout)

	if ok {
		patches, haveConflict := getPatchesFromInfo(info)
		if haveConflict {
			id := readMergeHead(repoPath)
			return WipPatches{Patches: patches, ConflictCommitId: &id}, nil
		}

		return WipPatches{Patches: patches}, nil
	}

	return WipPatches{}, nil
}

func getPatchesFromInfo(info []wipPatchInfo) (patches []WipPatch, haveConflict bool) {
	for _, i := range info {
		conflicted := isConflicted(i.staged, i.unStaged)
		if conflicted {
			haveConflict = true
		}
		patchType := pickTypeFromPatch(i.unStaged, i.staged)

		patches = append(
			patches, WipPatch{
				OldFile:      i.oldFile,
				NewFile:      i.newFile,
				PatchType:    patchType,
				StagedType:   i.staged,
				UnstagedType: i.unStaged,
				Conflicted:   conflicted,
				Id:           fmt.Sprintf("%v%v", i.newFile, patchType),
				IsImage:      false,
			},
		)
	}

	slices.SortFunc(
		patches, func(a, b WipPatch) int {
			return cmp.Compare(strings.ToLower(a.NewFile), strings.ToLower(b.NewFile))
		},
	)

	if haveConflict {
		var conflicted []WipPatch

		for _, patch := range patches {
			if patch.Conflicted {
				conflicted = append(conflicted, patch)
			}
		}
		return conflicted, haveConflict
	}

	return patches, haveConflict
}

func isConflicted(left wippatchtype.Type, right wippatchtype.Type) bool {
	return left == wippatchtype.U ||
		right == wippatchtype.U ||
		left == wippatchtype.A && right == wippatchtype.A ||
		left == wippatchtype.D && right == wippatchtype.D
}

func pickTypeFromPatch(unStaged wippatchtype.Type, staged wippatchtype.Type) wippatchtype.Type {
	if unStaged != wippatchtype.Empty {
		if unStaged == wippatchtype.Question {
			return wippatchtype.A
		}
		return unStaged
	}
	if staged == wippatchtype.Question {
		return wippatchtype.A
	}
	return staged
}

func readMergeHead(repoPath string) (commitId string) {
	rp, ok := cache.GetRepoPath(repoPath)
	if !ok {
		log.Warn("Called readMergeHead before the repoPath is initialised")
		return ""
	}
	bytes, err := os.ReadFile(path.Join(rp.GitPath, "MERGE_HEAD"))
	if err == nil {
		return string(bytes)
	}

	bytes, err = os.ReadFile(path.Join(rp.GitPath, "AUTO_MERGE"))
	if err == nil {
		return string(bytes)
	}
	return ""
}
