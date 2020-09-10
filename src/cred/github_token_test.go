package cred

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGithubToken(t *testing.T) {
	token := GetGithubToken()

	// TODO: Fill the second parameter "token"
	assert.EqualValues(t, token, "actual_token needs to be updated")
}
