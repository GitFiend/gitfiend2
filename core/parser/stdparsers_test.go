package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUntilLineEndParser(t *testing.T) {
	res, ok := ParsePart(UntilLineEnd, "asdfsdf&^HF JC\tasd !@\nasdf")

	assert.True(t, ok)
	assert.Equal(t, "asdfsdf&^HF JC\tasd !@", res)
}

func TestAnyWord(t *testing.T) {
	res, ok := ParsePart(AnyWord, "abcd55")

	assert.True(t, ok)
	assert.Equal(t, "abcd55", res)

	res, ok = ParsePart(AnyWord, "@@@@@")

	assert.False(t, ok)
	assert.Equal(t, "", res)
}

func TestSignedIntParser(t *testing.T) {
	res, ok := ParseAll(SignedInt, "1234")

	assert.True(t, ok)
	assert.Equal(t, "1234", res)
}
