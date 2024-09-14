package git

import (
	"gitfiend2/core/git/wippatchtype"
	p "gitfiend2/core/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPRenameStatus(t *testing.T) {
	out, ok := p.ParseAll(pRenameStatus, "R ")
	assert.True(t, ok)
	assert.Equal(t, 'R', out.R1)
	assert.Equal(t, ' ', out.R2)

	out, ok = p.ParseAll(pRenameStatus, " R")
	assert.True(t, ok)
	assert.Equal(t, ' ', out.R1)
	assert.Equal(t, 'R', out.R2)

	out, ok = p.ParseAll(pRenameStatus, "RM")
	assert.True(t, ok)
	assert.Equal(t, 'R', out.R1)
	assert.Equal(t, 'M', out.R2)

	out, ok = p.ParseAll(pRenameStatus, "DR")
	assert.True(t, ok)
	assert.Equal(t, 'D', out.R1)
	assert.Equal(t, 'R', out.R2)
}

func TestPWipRenamePatch(t *testing.T) {
	out, ok := p.ParseAll(pWipRenamePatch, "R  filename.txt\000has some spaces.txt\000")

	assert.True(t, ok)
	assert.Equal(
		t, wipPatchInfo{
			oldFile:  "has some spaces.txt",
			newFile:  "filename.txt",
			staged:   wippatchtype.R,
			unStaged: wippatchtype.Empty,
		},
		out,
	)
}

func TestPWorkStatus(t *testing.T) {
	out, _ := p.ParseAll(pWorkStatus, "??")
	assert.Equal(t, p.And2Result[rune, rune]{R1: '?', R2: '?'}, out)

	out, _ = p.ParseAll(pWorkStatus, " A")
	assert.Equal(t, p.And2Result[rune, rune]{R1: ' ', R2: 'A'}, out)

	out, _ = p.ParseAll(pWorkStatus, "AM")
	assert.Equal(t, p.And2Result[rune, rune]{R1: 'A', R2: 'M'}, out)
}

func TestPWipOtherPatch(t *testing.T) {
	out, ok := p.ParseAll(pWipOtherPatch, "DU folder/has a space/test2.js\000")
	assert.True(t, ok)
	assert.Equal(
		t, wipPatchInfo{
			oldFile:  "folder/has a space/test2.js",
			newFile:  "folder/has a space/test2.js",
			staged:   wippatchtype.D,
			unStaged: wippatchtype.U,
		}, out,
	)
}

func TestPWipPatches(t *testing.T) {
	out, ok := p.ParseAll(
		pWipPatches,
		"R  582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt\000 582160ee-5216-4dc6-bf74-1c1fce4978eb.txt\000DU folder/has a space/test2.js\000",
	)
	assert.True(t, ok)
	assert.Equal(
		t, []wipPatchInfo{
			{
				staged:   wippatchtype.R,
				unStaged: wippatchtype.Empty,
				oldFile:  "582160ee-5216-4dc6-bf74-1c1fce4978eb.txt",
				newFile:  "582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt",
			},
			{
				staged:   wippatchtype.D,
				unStaged: wippatchtype.U,
				oldFile:  "folder/has a space/test2.js",
				newFile:  "folder/has a space/test2.js",
			},
		}, out,
	)

	out, ok = p.ParseAll(
		pWipPatches,
		"C  582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt\000 582160ee-5216-4dc6-bf74-1c1fce4978eb.txt\000DU folder/has a space/test2.js\000",
	)
	assert.True(t, ok)
	assert.Equal(
		t, []wipPatchInfo{
			{
				staged:   wippatchtype.C,
				unStaged: wippatchtype.Empty,
				oldFile:  "582160ee-5216-4dc6-bf74-1c1fce4978eb.txt",
				newFile:  "582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt",
			},
			{
				staged:   wippatchtype.D,
				unStaged: wippatchtype.U,
				oldFile:  "folder/has a space/test2.js",
				newFile:  "folder/has a space/test2.js",
			},
		}, out,
	)
}

func TestPWipPatches2(t *testing.T) {
	out, ok := p.ParseAll(
		pWipPatches,
		"T  582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt\000DU folder/has a space/test2.js\000",
	)
	assert.True(t, ok)
	assert.Equal(
		t, []wipPatchInfo{
			{
				staged:   wippatchtype.T,
				unStaged: wippatchtype.Empty,
				oldFile:  "582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt",
				newFile:  "582160ee-5216-4dc6-bf74-1c1fce4978eb2.txt",
			},
			{
				staged:   wippatchtype.D,
				unStaged: wippatchtype.U,
				oldFile:  "folder/has a space/test2.js",
				newFile:  "folder/has a space/test2.js",
			},
		}, out,
	)
}

func TestPWipPatches3(t *testing.T) {
	text := " M .DS_Store\000 D LabBook/.ztr-directory\000 M LabBook/2023-06-18_CRISPR23-code.md\000?? Icon\r\000?? LabBook/2023-06-26_TEST.md\000"
	out, ok := p.ParseAll(pWipPatches, text)
	assert.True(t, ok)
	assert.Equal(
		t, []wipPatchInfo{
			{
				staged:   wippatchtype.Empty,
				unStaged: wippatchtype.M,
				oldFile:  ".DS_Store",
				newFile:  ".DS_Store",
			},
			{
				staged:   wippatchtype.Empty,
				unStaged: wippatchtype.D,
				oldFile:  "LabBook/.ztr-directory",
				newFile:  "LabBook/.ztr-directory",
			},
			{
				staged:   wippatchtype.Empty,
				unStaged: wippatchtype.M,
				oldFile:  "LabBook/2023-06-18_CRISPR23-code.md",
				newFile:  "LabBook/2023-06-18_CRISPR23-code.md",
			},
			{
				staged:   wippatchtype.Question,
				unStaged: wippatchtype.Question,
				oldFile:  "Icon\r",
				newFile:  "Icon\r",
			},
			{
				staged:   wippatchtype.Question,
				unStaged: wippatchtype.Question,
				oldFile:  "LabBook/2023-06-26_TEST.md",
				newFile:  "LabBook/2023-06-26_TEST.md",
			},
		}, out,
	)
}
