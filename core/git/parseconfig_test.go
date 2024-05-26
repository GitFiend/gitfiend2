package git

import (
	"gitfiend2/core/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPHeading(t *testing.T) {
	res, ok := parser.ParseAll(pHeading, "[core]")
	assert.True(t, ok)
	assert.Equal(t, "core", res.String())

	res, ok = parser.ParseAll(pHeading, "[remote \"origin\"]")
	assert.True(t, ok)
	assert.Equal(t, "remote.origin", res.String())

	res, ok = parser.ParseAll(pHeading, "[branch \"my-branch-name\"]")
	assert.True(t, ok)
	assert.Equal(t, "branch.my-branch-name", res.String())

	res, ok = parser.ParseAll(pHeading, "[branch \"feature/my-branch-name\"]")
	assert.True(t, ok)
	assert.Equal(t, "branch.feature/my-branch-name", res.String())
}

func TestParseSimpleConfig(t *testing.T) {
	text := `
[core]
	repositoryformatversion = 0
	filemode = true 
`

	res, ok := MakeConfigLog(text)
	assert.True(t, ok)
	assert.Equal(t, "core.repositoryformatversion=0\ncore.filemode=true", res)
}

func TestParseRandomComments(t *testing.T) {
	text := `
; Comment
[core]
	repositoryformatversion = 0
	filemode = true 
# hello
`
	res, ok := MakeConfigLog(text)
	assert.True(t, ok)
	assert.Equal(t, "core.repositoryformatversion=0\ncore.filemode=true", res)
}

func TestParseRealConfig(t *testing.T) {
	text := `[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = https://github.com/GitFiend/gitfiend-core.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "main"]
	remote = origin
	merge = refs/heads/main
[branch "ssr-code-viewer"]
	remote = origin
	merge = refs/heads/ssr-code-viewer
`
	res, ok := MakeConfigLog(text)
	assert.True(t, ok)
	assert.Equal(t, `core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true
remote.origin.url=https://github.com/GitFiend/gitfiend-core.git
remote.origin.fetch=+refs/heads/*:refs/remotes/origin/*
branch.main.remote=origin
branch.main.merge=refs/heads/main
branch.ssr-code-viewer.remote=origin
branch.ssr-code-viewer.merge=refs/heads/ssr-code-viewer`, res)
}

func TestParseLargerConfig(t *testing.T) {
	text := `[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
	ignorecase = true
	precomposeunicode = true
# Some comment.
[remote "origin"]
	url = https://github.com/GitFiend/git-fiend.git
	fetch = +refs/heads/*:refs/remotes/origin/*

; Some comment 2.
[branch "main"]
	remote = origin
	merge = refs/heads/main
[branch "cleanup"]
	remote = origin
	merge = refs/heads/cleanup
[branch "commit-switcher"]
	remote = origin
	merge = refs/heads/commit-switcher
[branch "server"]
	remote = origin
	merge = refs/heads/server
[branch "ws"]
	remote = origin
	merge = refs/heads/ws
[branch "alt-toolbar"]
	remote = origin
	merge = refs/heads/alt-toolbar
[branch "alt-ref-view"]
	remote = origin
	merge = refs/heads/alt-ref-view
[branch "image-conflicts"]
	remote = origin
	merge = refs/heads/image-conflicts
[branch "auto-complete"]
	remote = origin
	merge = refs/heads/auto-complete
[branch "mac-app"]
	remote = origin
	merge = refs/heads/mac-app
[branch "try-tauri"]
	remote = origin
	merge = refs/heads/try-tauri
[branch "split-view"]
	remote = origin
	merge = refs/heads/split-view
[branch "ssr-code-viewer"]
	remote = origin
	merge = refs/heads/ssr-code-viewer
`
	_, ok := MakeConfigLog(text)
	assert.True(t, ok)
}
