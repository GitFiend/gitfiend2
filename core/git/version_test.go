package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGitVersion(t *testing.T) {
	res := ParseGitVersion("git version 2.32.0")
	assert.Equal(t, VersionInfo{2, 32, 0}, res)

	res = ParseGitVersion("git version 2.0.0-alpha.1+")
	assert.Equal(t, VersionInfo{2, 0, 0}, res)

	res = ParseGitVersion("git version 2.32")
	assert.Equal(t, VersionInfo{2, 32, 0}, res)

	res = ParseGitVersion("git version 2.32.1 (Apple Git-133)")
	assert.Equal(t, VersionInfo{2, 32, 1}, res)

	res = ParseGitVersion("git version 2.37.3.windows.1")
	assert.Equal(t, VersionInfo{2, 37, 3}, res)
}
