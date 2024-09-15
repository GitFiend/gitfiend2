package git

import (
	"fmt"
	"gitfiend2/core/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

const p1 = "src2/parser-lib/input.ts"
const p2 = "src2/renderer-process/redux-cache/repo-state/commits/commits-reducer.test.ts"

func TestPRenamePatch(t *testing.T) {
	log := fmt.Sprintf("R100\000%s\000%s\000", p1, p2)
	res, ok := parser.ParseAll(pRenamePatch, log)

	assert.True(t, ok)
	assert.Equal(
		t, patchData{
			id:        fmt.Sprintf("%s-R100", p2),
			oldFile:   p1,
			newFile:   p2,
			patchType: 'R',
		}, res,
	)
}

func TestPCopyPatch(t *testing.T) {
	log := fmt.Sprintf("C100\000%s\000%s\000", p1, p2)
	res, ok := parser.ParseAll(pCopyPatch, log)

	assert.True(t, ok)
	assert.Equal(
		t, patchData{
			id:        fmt.Sprintf("%s-C100", p2),
			oldFile:   p1,
			newFile:   p2,
			patchType: 'C',
		}, res,
	)
}

func TestPOtherPatch(t *testing.T) {
	log := fmt.Sprintf("M\000%s\000", p2)
	res, ok := parser.ParseAll(pOtherPatch, log)

	assert.True(t, ok)
	assert.Equal(
		t, patchData{
			id:        fmt.Sprintf("%s-M", p2),
			oldFile:   p2,
			newFile:   p2,
			patchType: 'M',
		}, res,
	)
}
