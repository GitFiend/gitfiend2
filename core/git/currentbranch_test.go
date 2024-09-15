package git

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestGetRefNameFromPath(t *testing.T) {
	refName := getRefNameFromPath(path.Join("aa", "bb", "cc", "dd"), "")
	assert.Equal(t, "aa/bb/cc/dd", refName)

	refName = getRefNameFromPath(path.Join("aa", "bb", "cc", "dd"), "aa")
	assert.Equal(t, "bb/cc/dd", refName)
}
