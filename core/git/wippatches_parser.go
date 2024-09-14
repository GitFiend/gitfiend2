package git

import (
	"gitfiend2/core/git/wippatchtype"
	p "gitfiend2/core/parser"
)

var pWorkStatusPart = p.Or(
	p.Rune(' '),
	p.Rune('?'),
	p.Rune('A'),
	p.Rune('C'),
	p.Rune('D'),
	p.Rune('M'),
	p.Rune('R'),
	p.Rune('R'),
	p.Rune('U'),
	p.Rune('T'),
)

var pRenameStatus = p.Or(
	p.And2(p.Rune('R'), pWorkStatusPart),
	p.And2(pWorkStatusPart, p.Rune('R')),
)

var pWipRenamePatch = p.Map(p.And5(pRenameStatus, p.Ws, p.UntilNul, p.Ws, p.UntilNul), mapToInfo)

func mapToInfo(res p.And5Result[p.And2Result[rune, rune], string, string, string, string]) wipPatchInfo {
	return wipPatchInfo{
		oldFile:  res.R5,
		newFile:  res.R3,
		staged:   wippatchtype.Type(res.R1.R1),
		unStaged: wippatchtype.Type(res.R1.R2),
	}
}

var pWorkStatus = p.And2(pWorkStatusPart, pWorkStatusPart)

var pWipOtherPatch = p.Map(
	p.And3(pWorkStatus, p.Ws, p.UntilNul),
	func(res p.And3Result[p.And2Result[rune, rune], string, string]) wipPatchInfo {
		return wipPatchInfo{
			oldFile:  res.R3,
			newFile:  res.R3,
			staged:   wippatchtype.Type(res.R1.R1),
			unStaged: wippatchtype.Type(res.R1.R2),
		}
	},
)

var pCopyStatus = p.Or(
	p.And2(p.Rune('C'), pWorkStatusPart),
	p.And2(pWorkStatusPart, p.Rune('C')),
)

var pWipCopyPatch = p.Map(p.And5(pCopyStatus, p.Ws, p.UntilNul, p.Ws, p.UntilNul), mapToInfo)

var pWipPatches = p.Many(p.Or(pWipRenamePatch, pWipCopyPatch, pWipOtherPatch))
