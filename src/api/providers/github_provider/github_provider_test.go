package github_provider

import (
	"testing"

	"github.com/davetweetlive/golang-microservices/src/api/clients/restclient"
	"github.com/davetweetlive/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, header, "token abc123")
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.StartMockups()
	restclient.StopMockups()
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
}
