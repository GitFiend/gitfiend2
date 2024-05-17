package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberParser(t *testing.T) {
	res, _ := ParseAll(Number, "123 fd3s")

	if res != "123" {
		t.Error("Expected 123, got ", res)
	}

	res2, _ := ParseAll(Number, "-1.009")

	if res2 != "-1.009" {
		t.Error("Expected -1.009, got ", res)
	}
}

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
