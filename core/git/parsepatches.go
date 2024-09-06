package git

import (
	"fmt"
	p "gitfiend2/core/parser"
)

var pManyPatchesWithCommitIds = p.Many(pPatchesWithCommitId)

var pPatchesWithCommitId = p.And3(
	p.UntilParser(p.And2(p.Rune(','), p.Ws)),
	pPatches,
	p.Or(p.UntilNul, p.Ws),
)

type patchData struct {
	id, oldFile, newFile string
	patchType            PatchType
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
		patchType: PatchType(res.R1.R1),
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
	), func(res rune) PatchType {
		return PatchType(res)
	},
)

var pOtherPatch = p.Map(
	p.And3(pStatus, p.UntilNul, p.UntilNul),
	makeFromOther,
)

func makeFromOther(res p.And3Result[PatchType, string, string]) patchData {
	return patchData{
		id:        fmt.Sprintf("%s-%c", res.R3, res.R1),
		patchType: res.R1,
		oldFile:   res.R3,
		newFile:   res.R3,
	}
}
