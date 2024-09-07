package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSubmoduleFile(t *testing.T) {
	text := "gitdir: ../.git/modules/fiend-ui"
	p := parseSubmoduleFile(text)

	assert.Equal(t, "../.git/modules/fiend-ui", p)
}
