package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/davetweetlive/golang-microservices/src/api/clients/restclient"
	"github.com/davetweetlive/golang-microservices/src/api/domain/github"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, header, "token abc123")
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockUp(restclient.Mock{
		Url:        `https://api.github.com/user/repos`,
		HttpMethod: http.MethodPost,
		Response:   nil,
		Err:        errors.New("Invalid rest client response"),
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid rest client response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockUp(restclient.Mock{
		Url:        `https://api.github.com/user/repos`,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        `https://api.github.com/user/repos`,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid JSON response body", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        `https://api.github.com/user/repos`,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{ "message": "Requires authentication", "documentation_url": "https://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user" }`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

func TestCreateRepoInvalidSucessResponse(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        `https://api.github.com/user/repos`,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{ "id": "123" }`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Error unmarshaling create repo response", err.Message)
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        `https://api.github.com/user/repos`,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{ "id": 293069384, "name": "go-test", "full_name": "davetweetlive/go-test" }`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 293069384, response.Id)
	assert.EqualValues(t, "go-test", response.Name)
	assert.EqualValues(t, "davetweetlive/go-test", response.FullName)

}
