package git

import p "gitfiend2/core/parser"

//var pPatchesWithCommitId = p.And3(
//	p.UntilParser(p.And2(p.Rune(','), p.Ws)),
//	)

type patchData struct {
	id, oldFile, newFile string
	patchType            PatchType
}

var pRenamePatch = p.And4(
	p.And2(p.Rune('R'), p.UInt),
	p.UntilNul,
	p.UntilNul,
	p.UntilNul,
)
