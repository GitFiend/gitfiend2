package git

import (
	"fmt"
	"gitfiend2/core/git/patchtype"
	p "gitfiend2/core/parser"
	"slices"
	"strings"
)

var pManyPatchesWithCommitIds = p.Map(p.Many(pPatchesWithCommitId), patchesToMap)

func patchesToMap(res []p.And3Result[string, []patchData, string]) map[string][]Patch {
	m := map[string][]Patch{}

	for _, r := range res {
		commitId := r.R1
		data := r.R2
		m[commitId] = patchDataToPatches(commitId, data)
	}
	return m
}

func patchDataToPatches(commitId string, data []patchData) []Patch {
	var patches []Patch
	for _, pd := range data {
		patches = append(
			patches, Patch{
				CommitId:  commitId,
				OldFile:   pd.oldFile,
				NewFile:   pd.newFile,
				PatchType: pd.patchType,
				Id:        pd.id,
				IsImage:   fileIsImage(pd.newFile),
			},
		)
	}
	return patches
}

var imageExtensions = []string{
	".apng",
	".bmp",
	".gif",
	".ico",
	".cur",
	".jpg",
	".jpeg",
	".png",
	".svg",
	".webp",
}

func fileIsImage(filename string) bool {
	name := strings.ToLower(filename)
	return slices.ContainsFunc(
		imageExtensions, func(ext string) bool {
			return strings.Contains(name, ext)
		},
	)
}

var pPatchesWithCommitId = p.And3(
	p.UntilParser(p.And2(p.Rune(','), p.Ws)),
	pPatches,
	p.Or(p.UntilNul, p.Ws),
)

type patchData struct {
	id, oldFile, newFile string
	patchType            patchtype.Type
}

var pPatches = p.Many(p.Or(pRenamePatch, pCopyPatch, pOtherPatch))

var pRenamePatch = p.Map(
	p.And4(
		p.And2(p.Rune('R'), p.UInt),
		p.UntilNul,
		p.UntilNul,
		p.UntilNul,
	),
	makePatchData,
)

var pCopyPatch = p.Map(
	p.And4(
		p.And2(p.Rune('C'), p.UInt),
		p.UntilNul,
		p.UntilNul,
		p.UntilNul,
	),
	makePatchData,
)

func makePatchData(
	res p.And4Result[
		p.And2Result[rune, string],
		string,
		string,
		string],
) patchData {
	return patchData{
		id:        fmt.Sprintf("%s-%c%s", res.R4, res.R1.R1, res.R1.R2),
		patchType: patchtype.Type(res.R1.R1),
		oldFile:   res.R3,
		newFile:   res.R4,
	}
}

var pStatus = p.Map(
	p.Or(
		p.Rune('A'),
		p.Rune('B'),
		p.Rune('C'),
		p.Rune('D'),
		p.Rune('M'),
		p.Rune('T'),
		p.Rune('U'),
		p.Rune('X'),
	), func(res rune) patchtype.Type {
		return patchtype.Type(res)
	},
)

var pOtherPatch = p.Map(
	p.And3(pStatus, p.UntilNul, p.UntilNul),
	makeFromOther,
)

func makeFromOther(res p.And3Result[patchtype.Type, string, string]) patchData {
	return patchData{
		id:        fmt.Sprintf("%s-%c", res.R3, res.R1),
		patchType: res.R1,
		oldFile:   res.R3,
		newFile:   res.R3,
	}
}
