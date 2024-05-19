package server

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassData(t *testing.T) {
	data := []byte(`{"repoPath": "."}`)

	var req ReqOptions
	err := json.Unmarshal(data, &req)

	assert.Nil(t, err)
}
